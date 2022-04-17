package api2

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	BaseURL     string
	TokenID     string
	Secret      string
	TLSInsecure bool
	HTTPClient  *http.Client
}

func NewClient(baseURL, tokenID, secret string, tlsInsecure bool) *Client {
	return &Client{
		BaseURL:     baseURL + "/api2/json",
		TokenID:     tokenID,
		Secret:      secret,
		TLSInsecure: tlsInsecure,
	}
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Add("Authorization", fmt.Sprintf("PVEAPIToken=%s=%s", c.TokenID, c.Secret))
	client := &http.Client{}
	client.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: c.TLSInsecure},
	}
	resp, err := client.Do(req)
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
	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
