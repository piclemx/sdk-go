package discovery

const eventResouce = "/events"

// Adds event resource for searching multiple occurrences
func BuildEventSearchReq() *APIRequest {
	return baseAPIReq().withResource(eventResouce)
}

// Add event resource and id for getting specific occurrence details
func BuildGetEventDetReq(id string) *APIRequest {
	return baseAPIReq().withResource(eventResouce + "/" + id)
}

// Add event resource and id for getting specific occurrence images
func BuildGetEventImgReq(id string) *APIRequest {
	return baseAPIReq().withResource(eventResouce + "/" + id + "/images")

}
