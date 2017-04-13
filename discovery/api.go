package discovery

import (
	"log"
	"net/http"
	"io/ioutil"
)

type Api struct {
	key string
	conf Configuration
}

func NewApi(key string, conf Configuration) *Api {
	return &Api{key:key, conf:conf}
}

func (api *Api) EventsByKeyword(keyword string) (string, error) {
	params := map[string]string{"keyword":keyword}
	resp, err := api.getEvents(params)
	if err != nil {
		log.Println("EventsByKeyword: could not get events")
		return "", err
	}
	return resp, nil
}

func (api *Api )getEvents(params map[string]string) (string, error) {

	req, err := api.buildGetEventReq(params)
	if err != nil {
		log.Println("getEvents:", err)
		return "", err
	}

	resp, err := api.call(req)
	if err != nil {
		log.Println("getEvents:",err)
		return "", err
	}

	return resp, nil

}

func (api *Api) buildGetEventReq(params map[string]string) (*http.Request, error) {

	req, err := http.NewRequest("GET", api.conf.url+"/events.json", nil)
	if err != nil {
		log.Println("buildGetEventReq:",err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("apikey", api.key)
	for key, value := range params {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (api *Api) call(req *http.Request) (string, error) {

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("call:",err)
		return "", err
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	return string(body), nil
}