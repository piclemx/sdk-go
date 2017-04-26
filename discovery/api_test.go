package discovery

import (
	"fmt"
	"github.com/piclemx/sdk-go/discovery/domain"
	"github.com/piclemx/sdk-go/parameters"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const (
	okEventsResponse = `{ "_embedded": { "events": [ {"id":"1"} ] } }`
	okEventResponse  = `{"id":"1"}`
	errorResponse    = `{"fake error json string"}`
	validAPIKey      = "validApiKey"
	invalidAPIKey    = "invalidApiKey"
)

func buildTestServer(okResponse string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("apikey") == validAPIKey {
			fmt.Fprint(w, okResponse)

		} else {
			fmt.Fprint(w, errorResponse)
		}

	}))
	return ts
}

func buildTimeoutServer(okResponse string) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(w, okResponse)

	}))
	return ts
}

func TestBuildGetEventSearchSuccess(t *testing.T) {
	ts := buildTestServer(okEventsResponse)
	defer ts.Close()
	api := NewAPI(Configuration{key: validAPIKey, url: ts.URL})

	var resp domain.EventResponse
	err := api.Call(BuildEventSearchReq().WithParam(parameters.Keyword, "test"), &resp)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if resp.Embedded.Events[0].Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallApiWithTimeout(t *testing.T) {
	ts := buildTimeoutServer(okEventsResponse)
	defer ts.Close()
	api := NewAPI(Configuration{key: validAPIKey, url: ts.URL, timeout: 10 * time.Millisecond})

	var resp domain.EventResponse
	err := api.Call(BuildEventSearchReq().WithParam(parameters.Keyword, "test"), &resp)

	if err == nil || !strings.Contains(err.Error(), "Client.Timeout") {
		t.Errorf("should have timeout")
	}
}

func TestGetKey(t *testing.T) {
	key := "key"
	api := NewAPI(DefaultConfiguration().WithKey(key))

	if api.Key() != "" && api.Key() != key {
		t.Errorf("Should have the same key")
	}
}

func TestBuildGetEventDetailsSuccess(t *testing.T) {
	ts := buildTestServer(okEventResponse)
	defer ts.Close()
	api := NewAPI(Configuration{key: validAPIKey, url: ts.URL})

	var resp domain.Event
	err := api.Call(BuildGetEventDetReq("test"), &resp)
	if err != nil {
		t.Errorf("error:", err)
	}

	if resp.Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestBuildGetEventDetailsError(t *testing.T) {
	ts := buildTestServer(okEventResponse)
	defer ts.Close()
	api := NewAPI(Configuration{key: invalidAPIKey, url: ts.URL})

	var resp domain.Event
	err := api.Call(BuildGetEventDetReq("test"), &resp)

	if err == nil {
		t.Errorf("should receive error")
	}
}

func TestBuildGetEventImagesSuccess(t *testing.T) {
	ts := buildTestServer(okEventResponse)
	defer ts.Close()
	api := NewAPI(Configuration{key: validAPIKey, url: ts.URL})

	var resp domain.Event
	err := api.Call(BuildGetEventImgReq("test"), &resp)
	if err != nil {
		t.Errorf("error:", err)
	}

	if resp.Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestBuildGetEventImagesError(t *testing.T) {
	ts := buildTestServer(okEventResponse)
	defer ts.Close()
	api := NewAPI(Configuration{key: invalidAPIKey, url: ts.URL})

	var resp domain.Event
	err := api.Call(BuildGetEventImgReq("test"), &resp)

	if err == nil {
		t.Errorf("should receive error")
	}
}
