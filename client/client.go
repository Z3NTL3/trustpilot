package client

import (
	"net/http"
	"strings"

	"github.com/z3ntl3/trustpilot/model"
)

type Client struct {
	hello string
	*http.Client
}

// New creates a new instance of the Client struct.
//
// It initializes the Client struct with the default HTTP client provided by the net/http package.
// The function returns newly created Client.
func New() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

// SetHTTPClient sets the HTTP client for the Client struct.
//
// It takes a pointer to an http.Client as a parameter and updates the Client struct's Client field with the provided client.
// This function does not return anything.
func (c *Client) SetHTTPClient(client *http.Client) {
	c.Client = client
}

// SetTransport sets the HTTP transport for the Client struct.
//
// It takes a http.RoundTripper as a parameter and updates the Client struct's Transport field with the provided transport.
// This function does not return anything.
func (c *Client) SetTransport(transport http.RoundTripper) {
	c.Client.Transport = transport
}

// Execute sends an HTTP GET request to the specified URL and returns the parsed result.
//
// It takes a URL string as a parameter and checks if it contains "trustpilot.".
// If the URL does not contain "trustpilot.", it returns an error of type ErrNotAProperty.
//
// The function creates a new HTTP GET request using the provided URL and the HTTP client.
// It then sends the request and receives a response.
//
// The response is passed to the Parse function of the model package to parse the response body.
//
// Returns a pointer to a Result struct and an error.
func (c *Client) Execute(url string) (*model.Result, error) {
	if !strings.Contains(url, "trustpilot.") {
		return nil, ErrNotAProperty
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return model.Parse(resp.Body)
}