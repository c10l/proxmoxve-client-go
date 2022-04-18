package api2

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

type Client struct {
	BaseURL     string
	TokenID     string
	Secret      string
	TLSInsecure bool
	HTTPClient  *http.Client
}

type Response struct {
	Data any `json:"data"`
}

func NewClient(baseURL, tokenID, secret string, tlsInsecure bool) *Client {
	httpClient := new(http.Client)
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: tlsInsecure},
	}
	return &Client{
		BaseURL:     strings.Trim(baseURL, "/") + "/api2/json",
		TokenID:     tokenID,
		Secret:      secret,
		TLSInsecure: tlsInsecure,
		HTTPClient:  httpClient,
	}
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

func doGet[T any](c *Client, data *T, url string) (*T, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	response := new(Response)
	if err := json.Unmarshal(bytes, response); err != nil {
		return nil, err
	}
	responseData, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(responseData, &data); err != nil {
		return nil, err
	}
	return data, nil
}

func doPost[T any](c *Client, data *T, url string, body io.Reader) (*T, error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	bytes, err := c.doRequest(req)
	if err != nil {
		return nil, err
	}
	response := new(Response)
	if err := json.Unmarshal(bytes, response); err != nil {
		return nil, err
	}
	responseData, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(responseData, &data); err != nil {
		return nil, err
	}
	return data, nil
}
