package main

import (
	"fmt"
	"log"
	"os"

	"github.com/piclemx/sdk-go/discovery"
)

func main() {

	apikey := os.Getenv("TM_API_KEY")
	if apikey == "" {
		log.Println("can't find TM_API_KEY")
		return
	}

	api := discovery.NewApi(apikey, discovery.DefaultConfiguration())
	resp, err := api.EventsByKeyword("ed sheeran")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
}
