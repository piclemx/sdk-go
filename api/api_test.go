package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

type testStruct struct {
	Test string `json:"test"`
}

const (
	okResponse    = `{"test":"response"}`
	errorResponse = `{"fake error json string"}`
	validAPIKey   = "validApiKey"
	invalidAPIKey = "invalidApiKey"
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

func TestCallSuccess(t *testing.T) {
	ts := buildTestServer(okResponse)
	defer ts.Close()
	api := NewAPI(Configuration{Key: validAPIKey, URL: ts.URL})

	var resp testStruct
	err := api.Call(BaseAPIReq(), &resp)
	if err != nil {
		t.Errorf("error: %s", err)
	}

	if resp.Test != "response" {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallError(t *testing.T) {
	ts := buildTestServer(okResponse)
	defer ts.Close()
	api := NewAPI(Configuration{Key: invalidAPIKey, URL: ts.URL})

	var resp testStruct
	err := api.Call(BaseAPIReq(), &resp)

	if err == nil {
		t.Errorf("should receive error")
	}
}

func TestCallApiWithTimeout(t *testing.T) {
	ts := buildTimeoutServer(okResponse)
	defer ts.Close()
	api := NewAPI(Configuration{Key: validAPIKey, URL: ts.URL, Timeout: 10 * time.Millisecond})

	var resp testStruct
	err := api.Call(BaseAPIReq(), &resp)

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
