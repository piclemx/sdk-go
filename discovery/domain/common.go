package domain

type Link struct {
	Href      string `json:"href"`
	Templated bool   `json:"templated"`
}

type Links struct {
	Self Link `json:"self"`
	Next Link `json:"next"`
	Prev Link `json:"prev"`
}

type Loc struct {
	Longitude float64 `json:"longitute"`
	Latitude  float64 `json:"latitude"`
}

type Image struct {
	Url         string `json:"url"`
	Ratio       string `json:"ratio"`
	Width       int    `json:"width"`
	Height      int    `json:"height"`
	Fallback    bool   `json:"fallback"`
	Attribution string `json:"attribution"`
}
