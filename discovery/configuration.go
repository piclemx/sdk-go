package discovery

import "time"

// Configuration : Struct that contains elememts for the API client
type Configuration struct {
	key     string
	url     string
	timeout time.Duration
}

// DefaultConfiguration function gives defaults values
func DefaultConfiguration() Configuration {
	return Configuration{
		key:     "",
		url:     "https://app.ticketmaster.com/discovery/v2",
		timeout: 2 * time.Second}
}

// WithKey return the same Configuration with a new key
func (c *Configuration) WithKey(key string) *Configuration {
	c.key = key
	return c
}
