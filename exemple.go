package main

import (
	"fmt"
	"log"
	"os"

	"github.com/piclemx/sdk-go/discovery"
	"github.com/piclemx/sdk-go/discovery/domain"
	"github.com/piclemx/sdk-go/parameters"
)

func main() {

	apikey := os.Getenv("TM_API_KEY")
	if apikey == "" {
		log.Println("can't find TM_API_KEY")
		return
	}

	api := discovery.NewAPI(discovery.DefaultConfiguration().WithKey(apikey))

	eventSearchReq := discovery.BuildEventSearchReq().WithParam(parameters.Keyword, "ed sheeran")
	var eventSearchResp domain.EventResponse
	err := api.Call(eventSearchReq, &eventSearchResp)
	if err != nil {
		log.Println(err)
	}
	for _, event := range eventSearchResp.Embedded.Events {
		fmt.Println(event.Id, event.Name, event.URL)
	}

	eventDetailsReq := discovery.BuildGetEventDetReq(eventSearchResp.Embedded.Events[0].Id)
	var eventDetailsResp domain.Event
	err = api.Call(eventDetailsReq, &eventDetailsResp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(eventDetailsResp.Id, eventDetailsResp.Name, eventDetailsResp.URL)

	eventImagesReq := discovery.BuildGetEventImgReq(eventDetailsResp.Id)
	var eventImagesResp domain.Event
	err = api.Call(eventImagesReq, &eventImagesResp)
	if err != nil {
		log.Println(err)
	}
	for _, image := range eventImagesResp.Images {
		fmt.Println(image.Ratio, image.Url)
	}
}
