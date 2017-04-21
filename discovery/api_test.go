package discovery

import (
	"fmt"
	"github.com/piclemx/sdk-go/parameters"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

const (
	okResponse    = `{"fake ok response json string"}`
	errorResponse = `{"fake error json string"}`
	validAPIKey   = "validApiKey"
	invalidAPIKey = "invalidApiKey"
)

func buildTestServer() *httptest.Server {
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
	api := NewAPI(Configuration{key: validAPIKey, url: ts.URL})

	apiResp := api.Call(BuildEventSearchReq().WithParam(parameters.Keyword, "test"))

	resp := string(apiResp.Resp)
	if resp != okResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallApiWithError(t *testing.T) {
	ts := buildTestServer()
	defer ts.Close()
	api := NewAPI(Configuration{key: invalidAPIKey, url: ts.URL})

	apiResp := api.Call(BuildEventSearchReq().WithParam(parameters.Keyword, "test"))

	resp := string(apiResp.Resp)
	if resp != errorResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestCallApiWithTimeout(t *testing.T) {
	ts := buildTimeoutServer()
	defer ts.Close()
	api := NewAPI(Configuration{key: validAPIKey, url: ts.URL, timeout: 10 * time.Millisecond})

	apiResp := api.Call(BuildEventSearchReq().WithParam(parameters.Keyword, "test"))

	if apiResp.Err == nil || !strings.Contains(apiResp.Err.Error(), "Client.Timeout") {
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
	ts := buildTestServer()
	defer ts.Close()
	api := NewAPI(Configuration{key: validAPIKey, url: ts.URL})

	apiResp := api.Call(BuildEventSearchReq().WithParam(parameters.Keyword, "test"))

	resp := string(apiResp.Resp)
	if resp != okResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestBuildGetEventDetailsError(t *testing.T) {
	ts := buildTestServer()
	defer ts.Close()
	api := NewAPI(Configuration{key: invalidAPIKey, url: ts.URL})

	apiResp := api.Call(BuildGetEventDetReq("test"))

	resp := string(apiResp.Resp)
	if resp != errorResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestBuildGetEventImagesSuccess(t *testing.T) {
	ts := buildTestServer()
	defer ts.Close()
	api := NewAPI(Configuration{key: validAPIKey, url: ts.URL})

	apiResp := api.Call(BuildGetEventImgReq("test"))

	resp := string(apiResp.Resp)
	if resp != okResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}

func TestBuildGetEventImagesError(t *testing.T) {
	ts := buildTestServer()
	defer ts.Close()
	api := NewAPI(Configuration{key: invalidAPIKey, url: ts.URL})

	apiResp := api.Call(BuildGetEventImgReq("test"))

	resp := string(apiResp.Resp)
	if resp != errorResponse {
		t.Errorf("received incorrect response: %s", resp)
	}
}
