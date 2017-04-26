package main

import (
	"fmt"
	"log"
	"os"

	"github.com/piclemx/sdk-go/api"
	"github.com/piclemx/sdk-go/discovery"
	"github.com/piclemx/sdk-go/discovery/domain"
	"github.com/piclemx/sdk-go/discovery/parameters"
)

func main() {

	apikey := os.Getenv("TM_API_KEY")
	if apikey == "" {
		log.Println("can't find TM_API_KEY")
		return
	}

	api := api.NewAPI(api.DefaultConfiguration().WithKey(apikey))

	eventSearchResp := eventSearch(api)
	fmt.Println("Event search:")
	for _, event := range eventSearchResp.Embedded.Events {
		fmt.Println(event.Id, event.Name, event.URL)
	}

	eventDetailsResp := eventDetails(api, eventSearchResp.Embedded.Events[0].Id)
	fmt.Println("Event details:")
	fmt.Println(eventDetailsResp.Id, eventDetailsResp.Name, eventDetailsResp.URL)

	eventImagesResp := eventImages(api, eventDetailsResp.Id)
	fmt.Println("Event images:")
	for _, image := range eventImagesResp.Images {
		fmt.Println(image.Ratio, image.Url)
	}
}

func eventSearch(api *api.API) domain.EventResponse {
	eventSearchReq := discovery.BuildEventSearchReq().WithParam(parameters.Keyword, "ed sheeran")
	var eventSearchResp domain.EventResponse
	err := api.Call(eventSearchReq, &eventSearchResp)
	if err != nil {
		log.Println(err)
	}
	return eventSearchResp
}

func eventDetails(api *api.API, id string) domain.Event {
	eventDetailsReq := discovery.BuildGetEventDetReq(id)
	var eventDetailsResp domain.Event
	err := api.Call(eventDetailsReq, &eventDetailsResp)
	if err != nil {
		log.Println(err)
	}
	return eventDetailsResp
}

func eventImages(api *api.API, id string) domain.Event {
	eventImagesReq := discovery.BuildGetEventImgReq(id)
	var eventImagesResp domain.Event
	err := api.Call(eventImagesReq, &eventImagesResp)
	if err != nil {
		log.Println(err)
	}
	return eventImagesResp
}
