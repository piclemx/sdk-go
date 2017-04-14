package discovery

import "time"

type Configuration struct {
	url     string
	timeout time.Duration
}

func DefaultConfiguration() Configuration {
	return Configuration{
		url:     "https://app.ticketmaster.com/discovery/v2",
		timeout: 2 * time.Second}
}
