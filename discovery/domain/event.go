package domain

type Event struct {
	Type           string  `json:"type"`
	Distance       float64 `json:"distance"`
	Units          string  `json:"units"`
	Location       Loc     `json:"location"`
	Id             string  `json:"id"`
	Locale         string  `json:"locale"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	AdditionalInfo string  `json:"additionalinfo"`
	URL            string  `json:"url"`
	Images         []Image `json:"images"`
	Info           string  `json:"info"`
	PleaseNote     string  `json:"pleaseNote"`
	Test           bool    `json:"test"`
}

type EmbeddedEvents struct {
	Events []Event `json:"events"`
}

type EventResponse struct {
	Embedded EmbeddedEvents `json:"_embedded"`
}
