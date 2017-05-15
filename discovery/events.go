package discovery

import (
	"encoding/json"
	"github.com/piclemx/sdk-go/client"
	"github.com/piclemx/sdk-go/discovery/domain"
	"log"
)

const eventResource = "/events"

// Adds event resource for searching multiple occurrences
func BuildEventSearchReq() *client.APIRequest {
	return client.BaseAPIReq().WithResource(eventResource)
}

// Add event resource and id for getting specific occurrence details
func BuildGetEventDetReq(id string) *client.APIRequest {
	return client.BaseAPIReq().WithResource(eventResource + "/" + id)
}

// Add event resource and id for getting specific occurrence images
func BuildGetEventImgReq(id string) *client.APIRequest {
	return client.BaseAPIReq().WithResource(eventResource + "/" + id + "/images")
}

// Calls API to get domain.Events
func CallForEvents(client *client.Client, request *client.APIRequest) (*domain.Events, error) {
	resp, err := client.Call(request)
	if err != nil {
		log.Println("CallForEvents:", err.Error())
		return nil, err
	}

	var eventResponse domain.EventResponse
	err = json.Unmarshal([]byte(resp), &eventResponse)
	if err != nil {
		log.Println("CallForEvents:", err.Error())
		return nil, err
	}
	return &eventResponse.Embedded, nil
}

// Calls API to get domain.Event
func CallForEvent(client *client.Client, request *client.APIRequest) (*domain.Event, error) {
	resp, err := client.Call(request)
	if err != nil {
		log.Println("CallForEvents:", err.Error())
		return nil, err
	}

	var event domain.Event
	err = json.Unmarshal([]byte(resp), &event)
	if err != nil {
		log.Println("CallForEvents:", err.Error())
		return nil, err
	}
	return &event, nil

}
