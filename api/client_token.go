package api

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"strings"
)

type APIToken struct {
	TokenID string
	Secret  string
}

func NewAPITokenClient(baseURL, tokenID, secret string, tlsInsecure bool) (*Client, error) {
	httpClient := new(http.Client)
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: tlsInsecure},
	}
	apiURL, err := url.Parse(strings.TrimRight(baseURL, "/"))
	if err != nil {
		return nil, err
	}
	apiURL.Path += "/api2/json"

	client := &Client{
		APIurl:      *apiURL,
		APIToken:    &APIToken{TokenID: tokenID, Secret: secret},
		TLSInsecure: tlsInsecure,
		HTTPClient:  httpClient,
	}

	return client, nil
}
