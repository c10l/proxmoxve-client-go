package api2

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	ApiURL      *url.URL
	TokenID     string
	Secret      string
	TLSInsecure bool
	HTTPClient  *http.Client
}

func NewClient(baseURL, tokenID, secret string, tlsInsecure bool) (*Client, error) {
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
		ApiURL:      apiURL,
		TokenID:     tokenID,
		Secret:      secret,
		TLSInsecure: tlsInsecure,
		HTTPClient:  httpClient,
	}

	if _, err := client.RetrieveVersion(); err != nil {
		return nil, err
	}

	return client, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("Authorization", fmt.Sprintf("PVEAPIToken=%s=%s", c.TokenID, c.Secret))
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("%s\n%s", resp.Status, body)
	}
	return body, nil
}

func callAPI(c *Client, method string, url *url.URL) ([]byte, error) {
	req, err := http.NewRequest(method, url.String(), nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func doGet(c *Client, url *url.URL) ([]byte, error) {
	resp, err := callAPI(c, http.MethodGet, url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func doPost(c *Client, url *url.URL) ([]byte, error) {
	resp, err := callAPI(c, http.MethodPost, url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func doDelete(c *Client, url *url.URL) ([]byte, error) {
	return callAPI(c, http.MethodDelete, url)
}

func doPut(c *Client, url *url.URL) ([]byte, error) {
	return callAPI(c, http.MethodPut, url)
}
