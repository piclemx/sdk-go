package parameters

const (
	// Sorting order of the search result. eg.: relevance, desc
	Sort = "sort"
	// Filter events by latitude and longitude
	Latlong = "latlong"
	// Radius of the area in which we want to search for events.
	Radius = "radius"
	// Unit of the radius. [miles, km]
	Unit = "unit"
	// Filter events with a start date after this date
	StartDateTime = "startDateTime"
	// Filter events with a start date before this date
	EndDateTime = "endDateTime"
	// Filter events with onsale start date after this date
	OnsaleStartDateTime = "onsaleStartDateTime"
	// Filter events with onsale end date before this date
	OnsaleEndDateTime = "onsaleEndDateTime"
	// Filter events by country code
	CountryCode = "countryCode"
	// Filter events by state code
	StateCode = "stateCode"
	// Filter events by venue id
	VenueId = "venueId"
	// Filter events by attraction id
	AttractionId = "attractionId"
	// Filter events by segment id
	SegmentId = "segmentId"
	// Filter events by segment name
	SegmentName = "segmentName"
	// Filter events by classification name: name of any segment, genre, sub-genre, type, sub-type. [array]
	ClassificationName = "classificationName"
	// Filter events by classification id: id of any segment, genre, sub-genre, type, sub-type. [array]
	ClassificationId = "classificationId"
	// Filter events by market id
	MarketId = "marketId"
	// Filter events by promoter id
	PromoterId = "promoterId"
	// Filter events by dma id
	DmaId = "dmaId"
	// True, to include events with date to be announce (TBA). [yes, no, only]
	IncludeTBA = "includeTBA"
	// True, to include event with a date to be defined (TBD). [yes, no, only]
	IncludeTBD = "includeTBD"
	// Filter events by clientName
	ClientVisibility = "clientVisibility"
	// Keyword to search on
	Keyword = "keyword"
	// Filter entities by its id
	Id = "id"
	// Filter entities by its source name. [ticketmaster, universe, frontgate, tmr]
	Source = "source"
	// True if you want to have entities flag as test in the response. Only, if you only wanted test entities. [yes, no, only]
	IncludeTest = "includeTest"
	// Page number
	Page = "page"
	// Page size of the response
	Size = "size"
	// The locale in ISO code format. Multiple comma-separated values can be provided. When omitting the country part of the code (e.g. only 'en' or 'fr') then the first matching locale is used.
	Locale = "locale"
)
