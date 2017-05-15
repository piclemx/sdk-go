package client

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Client struct contains the http client and its configuration.
type Client struct {
	conf   Configuration
	client *http.Client
}

// NewClient : Creation of a new client
func NewClient(conf Configuration) *Client {
	return &Client{conf: conf}
}

// Key return the current Key for the api.
func (client *Client) Key() string {
	return client.conf.Key
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

// Call discovery API to get json string or an error
func (client *Client) Call(apiReq *APIRequest) (string, error) {
	if client.client == nil {
		client.client = &http.Client{Timeout: client.conf.Timeout}
	}

	apiReq.WithParam("apikey", client.Key())
	req, err := client.buildHttpReq(apiReq)
	if err != nil {
		log.Println("call:", err)
		return "", err
	}

	log.Println(req.URL.String())
	log.Flags()
	resp, err := client.client.Do(req)
	if err != nil {
		log.Println("call:", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("call:", err)
		return "", err
	}

	return string(body), nil
}

// Creates a basic API request
func BaseAPIReq() *APIRequest {
	apiReq := &APIRequest{method: "GET", params: make(map[string]string)}
	return apiReq
}

// Add resource to the API request
func (apiRep *APIRequest) WithResource(resource string) *APIRequest {
	apiRep.resource = resource
	return apiRep
}

func (client *Client) buildHttpReq(request *APIRequest) (*http.Request, error) {

	req, err := http.NewRequest(request.method, client.urlWith(request.resource), nil)
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

func (client *Client) urlWith(resource string) string {
	return client.conf.URL + resource
}
