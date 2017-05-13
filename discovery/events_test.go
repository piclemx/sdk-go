package discovery

import (
	"fmt"
	"github.com/piclemx/sdk-go/client"
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

func TestCallForEventsSearch(t *testing.T) {
	ts := buildTestServer(okEventsResponse)
	defer ts.Close()
	client := client.NewClient(client.Configuration{Key: validAPIKey, URL: ts.URL})

	resp, err := CallForEvents(client, BuildEventSearchReq().WithParam(parameters.Keyword, "test"))
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if resp.Events[0].Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallForEventDetail(t *testing.T) {
	ts := buildTestServer(okEventResponse)
	defer ts.Close()
	client := client.NewClient(client.Configuration{Key: validAPIKey, URL: ts.URL})

	resp, err := CallForEvent(client, BuildGetEventDetReq("test"))
	if err != nil {
		t.Errorf("error:", err)
	}

	if resp.Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallForEventImages(t *testing.T) {
	ts := buildTestServer(okEventResponse)
	defer ts.Close()
	client := client.NewClient(client.Configuration{Key: validAPIKey, URL: ts.URL})

	resp, err := CallForEvent(client, BuildGetEventImgReq("test"))
	if err != nil {
		t.Errorf("error:", err)
	}

	if resp.Id != "1" {
		t.Errorf("received incorrect response: %s", resp)
	}
}
