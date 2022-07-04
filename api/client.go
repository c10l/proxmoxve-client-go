package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Client struct {
	ApiURL      *url.URL
	APIToken    *APIToken
	UserPass    *UserPass
	TLSInsecure bool
	HTTPClient  *http.Client
}

type APIToken struct {
	TokenID string
	Secret  string
}

type UserPass struct {
	User   string
	Pass   string
	Ticket Ticket
}

type Ticket struct {
	CSRFPreventionToken string `json:"CSRFPreventionToken"`
	Ticket              string `json:"ticket"`
	Username            string `json:"username"`
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
		ApiURL:      apiURL,
		APIToken:    &APIToken{TokenID: tokenID, Secret: secret},
		TLSInsecure: tlsInsecure,
		HTTPClient:  httpClient,
	}

	return client, nil
}

func NewTicketClient(baseURL, user, pass string, tlsInsecure bool) (*Client, error) {
	httpClient := new(http.Client)
	httpClient.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: tlsInsecure},
	}
	apiURL, err := url.Parse(strings.TrimRight(baseURL, "/"))
	if err != nil {
		return nil, err
	}
	apiURL.Path += "/api2/json"

	// Retrieve ticket
	ticketURL := *apiURL
	ticketURL.Path += "/access/ticket"
	data := url.Values{
		"username": {user},
		"password": {url.QueryEscape(pass)},
	}
	ticketResp, err := httpClient.PostForm(ticketURL.String(), data)
	if err != nil {
		return nil, err
	}
	defer ticketResp.Body.Close()
	ticketBody, err := ioutil.ReadAll(ticketResp.Body)
	if err != nil {
		return nil, err
	}
	if ticketResp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("%s: %s", ticketResp.Status, ticketBody)
	}
	ticketRaw, err := parseData(ticketBody)
	if err != nil {
		return nil, err
	}
	ticket := Ticket{}
	if err := json.Unmarshal(ticketRaw, &ticket); err != nil {
		return nil, err
	}

	client := &Client{
		ApiURL: apiURL,
		UserPass: &UserPass{
			User:   user,
			Pass:   pass,
			Ticket: ticket,
		},
		TLSInsecure: tlsInsecure,
		HTTPClient:  httpClient,
	}

	return client, nil
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.UserPass != nil {
		req.Header.Set("CSRFPreventionToken", c.UserPass.Ticket.CSRFPreventionToken)
		req.Header.Set("Cookie", "PVEAuthCookie="+c.UserPass.Ticket.Ticket)
		req.Header.Set("Accept", "application/json")
	}
	if c.APIToken != nil {
		req.Header.Add("Authorization", fmt.Sprintf("PVEAPIToken=%s=%s", c.APIToken.TokenID, c.APIToken.Secret))
	}
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
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
	return parseData(resp)
}

func parseData(resp []byte) ([]byte, error) {
	type responseType struct {
		Data any `json:"data,omitempty"`
	}
	response := responseType{}
	if err := json.Unmarshal(resp, &response); err != nil {
		return nil, err
	}
	return json.Marshal(response.Data)
}

func (c *Client) Get(url *url.URL) ([]byte, error) {
	return callAPI(c, http.MethodGet, url)
}

func (c *Client) Post(url *url.URL) ([]byte, error) {
	return callAPI(c, http.MethodPost, url)
}

func (c *Client) Delete(url *url.URL) ([]byte, error) {
	return callAPI(c, http.MethodDelete, url)
}

func (c *Client) Put(url *url.URL) ([]byte, error) {
	return callAPI(c, http.MethodPut, url)
}
