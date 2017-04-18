package main

import (
	"fmt"
	"log"
	"os"

	"github.com/piclemx/sdk-go/discovery"
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
	eventSearchResp, err := api.Call(eventSearchReq)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(eventSearchResp)

	eventDetailsReq := discovery.BuildGetEventDetReq("16vZZfJ-wG7nv60")
	eventDetailsResp, err := api.Call(eventDetailsReq)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(eventDetailsResp)

	eventImagesReq := discovery.BuildGetEventImgReq("16vZZfJ-wG7nv60")
	eventImagesResp, err := api.Call(eventImagesReq)
	if err != nil {
		log.Println(err)
	}
	fmt.Println(eventImagesResp)

}
