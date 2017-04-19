package main

import (
	"fmt"
	"log"
	"os"

	"github.com/piclemx/sdk-go/discovery"
	"github.com/piclemx/sdk-go/parameters"
	"github.com/piclemx/sdk-go/discovery/domain"
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
	err := api.Call(eventSearchReq).WriteTo(&eventSearchResp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(eventSearchResp.Embedded.Events)

	eventDetailsReq := discovery.BuildGetEventDetReq(eventSearchResp.Embedded.Events[0].Id)
	var eventDetailsResp domain.Event
	err = api.Call(eventDetailsReq).WriteTo(&eventDetailsResp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(eventDetailsResp)

	eventImagesReq := discovery.BuildGetEventImgReq(eventDetailsResp.Id)
	var eventImagesResp domain.Event
	err = api.Call(eventImagesReq).WriteTo(&eventImagesResp)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(eventImagesResp)
}
