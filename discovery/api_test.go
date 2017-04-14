package discovery

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const (
	okResponse    = `{"fake ok response json string"}`
	errorResponse = `{"fake error json string"}`
	validApiKey   = "validApiKey"
	invalidApiKey = "invalidApiKey"
)

func buildTestServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Query().Get("apikey") == validApiKey {
			fmt.Fprint(w, okResponse)
		} else {
			fmt.Fprint(w, errorResponse)
		}

	}))
	return ts
}

func buildTimeoutServer() *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		time.Sleep(100 * time.Millisecond)
		fmt.Fprint(w, okResponse)

	}))
	return ts
}

func TestCallApiWithSuccess(t *testing.T) {
	ts := buildTestServer()
	defer ts.Close()
	api := NewApi(validApiKey, Configuration{url: ts.URL})

	resp, _ := api.EventsByKeyword("test")

	if resp != okResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallApiWithError(t *testing.T) {
	ts := buildTestServer()
	defer ts.Close()
	api := NewApi(invalidApiKey, Configuration{url: ts.URL})

	resp, _ := api.EventsByKeyword("test")

	if resp != errorResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallApiWithTimeout(t *testing.T) {
	ts := buildTimeoutServer()
	defer ts.Close()
	api := NewApi(validApiKey, Configuration{url: ts.URL, timeout: 10 * time.Millisecond})

	_, err := api.EventsByKeyword("test")

	if err == nil || !strings.Contains(err.Error(), "Client.Timeout") {
		t.Errorf("should have timout")
	}
}
