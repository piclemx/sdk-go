package client

import "time"

// Configuration : Struct that contains elements for the client
type Configuration struct {
	Key     string
	URL     string
	Timeout time.Duration
}

// DefaultURL is the current api URL if none is provided
const DefaultURL string = "https://app.ticketmaster.com/discovery/v2"

// DefaultTimeout set to 2 seconds
const DefaultTimeout time.Duration = 2 * time.Second

// DefaultConfiguration function gives defaults values
func DefaultConfiguration() Configuration {
	return Configuration{
		Key:     "",
		URL:     DefaultURL,
		Timeout: DefaultTimeout}
}

// WithKey return the same Configuration with a new Key
func (c Configuration) WithKey(key string) Configuration {
	c.Key = key
	return c
}
