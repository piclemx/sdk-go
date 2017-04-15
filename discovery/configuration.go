package discovery

import "time"

// Configuration : Struct that contains elememts for the API client
type Configuration struct {
	key     string
	url     string
	timeout time.Duration
}

// DefaultURL is the current API URL if none is provided
const DefaultURL string = "https://app.ticketmaster.com/discovery/v2"

// DefaultTimeout set to 2 seconds
const DefaultTimeout time.Duration = 2 * time.Second

// DefaultConfiguration function gives defaults values
func DefaultConfiguration() Configuration {
	return Configuration{
		key:     "",
		url:     DefaultURL,
		timeout: DefaultTimeout}
}

// WithKey return the same Configuration with a new key
func (c Configuration) WithKey(key string) Configuration {
	c.key = key
	return c
}
