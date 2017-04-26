package discovery

import "github.com/piclemx/sdk-go/api"

const eventResource = "/events"

// Adds event resource for searching multiple occurrences
func BuildEventSearchReq() *api.APIRequest {
	return api.BaseAPIReq().WithResource(eventResource)
}

// Add event resource and id for getting specific occurrence details
func BuildGetEventDetReq(id string) *api.APIRequest {
	return api.BaseAPIReq().WithResource(eventResource + "/" + id)
}

// Add event resource and id for getting specific occurrence images
func BuildGetEventImgReq(id string) *api.APIRequest {
	return api.BaseAPIReq().WithResource(eventResource + "/" + id + "/images")
}
