package discovery

import (
	"fmt"
	"github.com/piclemx/sdk-go/api"
	"github.com/piclemx/sdk-go/discovery/domain"
	"github.com/piclemx/sdk-go/discovery/parameters"
	"net/http"
	"net/http/httptest"
	"testing"
)

const (
	okEventsResponse = `{ "_embedded": { "events": [ {"id":"1"} ] } }`
	okEventResponse  = `{"id":"1"}`
	errorResponse    = `{"fake error json string"}`
	validAPIKey      = "validApiKey"
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

func TestBuildGetEventSearch(t *testing.T) {
	ts := buildTestServer(okEventsResponse)
	defer ts.Close()
	api := api.NewAPI(api.Configuration{Key: validAPIKey, URL: ts.URL})

	var resp domain.EventResponse
	err := api.Call(BuildEventSearchReq().WithParam(parameters.Keyword, "test"), &resp)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if resp.Embedded.Events[0].Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestBuildGetEventDetails(t *testing.T) {
	ts := buildTestServer(okEventResponse)
	defer ts.Close()
	api := api.NewAPI(api.Configuration{Key: validAPIKey, URL: ts.URL})

	var resp domain.Event
	err := api.Call(BuildGetEventDetReq("test"), &resp)
	if err != nil {
		t.Errorf("error:", err)
	}

	if resp.Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestBuildGetEventImages(t *testing.T) {
	ts := buildTestServer(okEventResponse)
	defer ts.Close()
	api := api.NewAPI(api.Configuration{Key: validAPIKey, URL: ts.URL})

	var resp domain.Event
	err := api.Call(BuildGetEventImgReq("test"), &resp)
	if err != nil {
		t.Errorf("error:", err)
	}

	if resp.Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}
