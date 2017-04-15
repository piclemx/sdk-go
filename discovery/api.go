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

// EventsByKeyword : Get envts by keyword
func (api *API) EventsByKeyword(keyword string) (string, error) {
	params := map[string]string{"keyword": keyword}
	resp, err := api.getEvents(params)
	if err != nil {
		log.Println("EventsByKeyword: could not get events")
		return "", err
	}
	return resp, nil
}

func (api *API) getEvents(params map[string]string) (string, error) {

	req, err := api.buildGetEventReq(params)
	if err != nil {
		log.Println("getEvents:", err)
		return "", err
	}

	resp, err := api.call(req)
	if err != nil {
		log.Println("getEvents:", err)
		return "", err
	}

	return resp, nil

}

func (api *API) buildGetEventReq(params map[string]string) (*http.Request, error) {

	req, err := http.NewRequest("GET", api.conf.url+"/events.json", nil)
	if err != nil {
		log.Println("buildGetEventReq:", err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("apikey", api.conf.key)
	for key, value := range params {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (api *API) call(req *http.Request) (string, error) {

	if api.client == nil {
		api.client = &http.Client{Timeout: api.conf.timeout}
	}

	resp, err := api.client.Do(req)
	if err != nil {
		log.Println("call:", err)
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}
