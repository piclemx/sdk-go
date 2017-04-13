package discovery

type Configuration struct {
	url string

}

func DefaultConfiguration() Configuration{
	return Configuration{url:"https://app.ticketmaster.com/discovery/v2"}
}