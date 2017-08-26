package shopgun

import (
	"fmt"
	"log"
	"net/http"
)

// Client is a shopgun client
type Client struct {
	basePath string
	c        *http.Client
	apiKey   string
	secret   string
	sr       SessionResponse
	Log      log.Logger
}

// Pagination is what it says
type Pagination struct {
	Offset uint64
	Limit  uint64
}

// DefaultPagination returns a default pagination setting
func DefaultPagination() Pagination {
	return Pagination{Limit: 24, Offset: 0}
}

// New returns a new shopgun client
func New(apiKey, secret string) *Client {
	if apiKey == "" {
		panic("Api key not set")
	}
	if secret == "" {
		panic("Secret not set")
	}
	client := &http.Client{}

	return &Client{
		basePath: "https://api.etilbudsavis.dk/v2",
		c:        client,
		apiKey:   apiKey,
		secret:   secret,
	}
}

// DealerSearch performs a search for dealers and returns a list of matches
func (c *Client) DealerSearch(query string, pagination Pagination) ([]*Dealer, error) {
	p := map[string]string{
		"query":  query,
		"limit":  fmt.Sprint(pagination.Limit),
		"offset": fmt.Sprint(pagination.Offset),
	}
	path := buildPath("/dealers/search", p)

	dealers := []*Dealer{}
	err := c.doGet(path, dealers)
	if err != nil {
		return nil, err
	}
	return dealers, nil
}

// Catalogs returns a list of catalogs from dealer id's given
func (c *Client) Catalogs(dealerIDs []string, pagination Pagination) (*Catalogs, error) {
	p := map[string]string{
		"dealer_ids": stringListArgs(dealerIDs),
		"limit":      fmt.Sprint(pagination.Limit),
		"offset":     fmt.Sprint(pagination.Offset),
	}
	path := buildPath("/catalogs", p)

	cats := &Catalogs{}
	err := c.doGet(path, cats)
	if err != nil {
		return nil, err
	}
	return cats, nil
}

// Offers returns a list of offers from catalog id's given
func (c *Client) Offers(catalogIDs []string, pagination Pagination) (*Offers, error) {
	p := map[string]string{
		"catalog_ids": stringListArgs(catalogIDs),
		"limit":       fmt.Sprint(pagination.Limit),
		"offset":      fmt.Sprint(pagination.Offset),
	}
	path := buildPath("/offers", p)

	offers := &Offers{}
	err := c.doGet(path, offers)
	if err != nil {
		return nil, err
	}
	return offers, nil
}
