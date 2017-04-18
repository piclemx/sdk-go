package discovery

import (
	"io/ioutil"
	"log"
	"net/http"
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

// Call discovery API
func (api *API) Call(apiReq *APIRequest) (string, error) {
	if api.client == nil {
		api.client = &http.Client{Timeout: api.conf.timeout}
	}

	apiReq.WithParam("apikey", api.Key())
	req, err := api.buildHttpReq(apiReq)
	if err != nil {
		log.Println("call:", err)
		return "", err
	}

	log.Println(req.URL.String())
	resp, err := api.client.Do(req)
	if err != nil {
		log.Println("call:", err)
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
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