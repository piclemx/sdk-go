package discovery

import (
	"io/ioutil"
	"log"
	"net/http"
	"encoding/json"
)

// API struct contains the API client and is configuration.
type API struct {
	conf   Configuration
	client *http.Client
}

// NewAPI : Creation of a new client
func NewAPI(conf Configuration) *API {
	return &API{conf: conf}
}

// Key return the current key for the API.
func (api *API) Key() string {
	return api.conf.key
}

// API request method, resource and params
type APIRequest struct {
	method   string
	resource string
	params   map[string]string
}

// Adds parameters in the request
func (apiReq *APIRequest) WithParam(param string, value string) *APIRequest {
	apiReq.params[param] = value
	return apiReq
}

type ApiResponse struct {
	Resp []byte
	Err error
}

// Call discovery API
func (api *API) Call(apiReq *APIRequest) *ApiResponse {
	if api.client == nil {
		api.client = &http.Client{Timeout: api.conf.timeout}
	}

	apiReq.WithParam("apikey", api.Key())
	req, err := api.buildHttpReq(apiReq)
	if err != nil {
		log.Println("call:", err)
		return &ApiResponse{Err: err}
	}

	log.Println(req.URL.String())
	log.Flags()
	resp, err := api.client.Do(req)
	if err != nil {
		log.Println("call:", err)
		return &ApiResponse{Err: err}
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return &ApiResponse{body, err}
}

// Writes the response body to the variable passed
func (apiResp *ApiResponse) WriteTo(value interface{}) error {
	if apiResp.Err != nil {
		return apiResp.Err
	}

	err := json.Unmarshal(apiResp.Resp, &value)
	if err != nil {
		return err
	}

	return nil
}

func baseAPIReq() *APIRequest {
	apiReq := &APIRequest{method: "GET", params: make(map[string]string)}
	return apiReq
}

func (apiRep *APIRequest) withResource(resource string) *APIRequest {
	apiRep.resource = resource
	return apiRep
}

func (api *API) buildHttpReq(request *APIRequest) (*http.Request, error) {

	req, err := http.NewRequest(request.method, api.conf.url+request.resource, nil)
	if err != nil {
		log.Println("buildHttpReq:", err)
		return nil, err
	}

	q := req.URL.Query()
	for key, value := range request.params {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()
	return req, nil
}