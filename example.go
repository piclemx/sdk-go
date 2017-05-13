package main

import (
	"fmt"
	"log"
	"os"

	"github.com/piclemx/sdk-go/client"
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

	client := client.NewClient(client.DefaultConfiguration().WithKey(apikey))

	eventSearchResp := eventSearch(client)
	fmt.Println("Event search:")
	for _, event := range eventSearchResp.Events {
		fmt.Println(event.Id, event.Name, event.URL)
	}

	eventDetailsResp := eventDetails(client, eventSearchResp.Events[0].Id)
	fmt.Println("Event details:")
	fmt.Println(eventDetailsResp.Id, eventDetailsResp.Name, eventDetailsResp.URL)

	eventImagesResp := eventImages(client, eventDetailsResp.Id)
	fmt.Println("Event images:")
	for _, image := range eventImagesResp.Images {
		fmt.Println(image.Ratio, image.Url)
	}
}

func eventSearch(client *client.Client) *domain.Events {
	eventSearchReq := discovery.BuildEventSearchReq().WithParam(parameters.Keyword, "ed sheeran")
	resp, err := discovery.CallForEvents(client, eventSearchReq)
	if err != nil {
		log.Println(err)
	}
	return resp
}

func eventDetails(client *client.Client, id string) *domain.Event {
	eventDetailsReq := discovery.BuildGetEventDetReq(id)
	resp, err := discovery.CallForEvent(client, eventDetailsReq)
	if err != nil {
		log.Println(err)
	}
	return resp
}

func eventImages(client *client.Client, id string) *domain.Event {
	eventImagesReq := discovery.BuildGetEventImgReq(id)
	resp, err := discovery.CallForEvent(client, eventImagesReq)
	if err != nil {
		log.Println(err)
	}
	return resp
}
