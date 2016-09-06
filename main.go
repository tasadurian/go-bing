package Bing

import (
	"encoding/json"
	"errors"
	"net/http"
)

const (
	searchAPI    = "https://api.datamarket.azure.com/Bing/Search/"
	webSearchAPI = "https://api.datamarket.azure.com/Bing/SearchWeb/"
	// XML specifies parameter for XML results.
	XML = "ATOM"
	// JSON specifies parameter for JSON results.
	JSON = "JSON"
)

// Client is the structure used to initialize a new
// go-bing client. Needs a valid api key.
type Client struct {
	// API key
	key string
	// format either XML or JSON.
	format string
}

// SearchWeb is used to do a basic web search.
func (c *Client) SearchWeb(query string) {
	res, err := http.Get(query)
	if err != nil {
		panic(err)
	}
}
