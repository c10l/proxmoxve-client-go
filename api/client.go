package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Client struct {
	APIurl   url.URL
	APIToken *APIToken
	// UserPass    *UserPass
	Ticket      *Ticket
	TLSInsecure bool
	HTTPClient  *http.Client
}

func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	if c.Ticket != nil {
		req.Header.Set("CSRFPreventionToken", c.Ticket.CSRFPreventionToken)
		req.Header.Set("Cookie", "PVEAuthCookie="+c.Ticket.Ticket)
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
	body, err := io.ReadAll(resp.Body)
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

type ItemDeleter interface {
	Delete() error
}

func (c *Client) DeleteItem(item ItemDeleter, basePath, id string, digest string) error {
	if id == "" {
		return fmt.Errorf("Client.DeleteItem: item ID is required")
	}

	apiURL := c.APIurl
	apiURL.Path += basePath + "/" + id

	params := url.Values{}
	if digest != "" {
		params.Add("digest", digest)
	}
	apiURL.RawQuery = params.Encode()

	_, err := c.Delete(&apiURL)
	return err
}

type ItemGetter interface {
	GetItem() ([]byte, error)
}

func (c Client) GetItem(g ItemGetter, basePath, id string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("Client.GetItem: item ID is required")
	}

	apiURL := c.APIurl
	apiURL.Path += basePath + "/" + id
	return c.Get(&apiURL)
}

type ItemPutter interface {
	PutItem() ([]byte, error)
	ParseParams(*url.URL) error
}

func (c Client) PutItem(p ItemPutter, basePath, id string) ([]byte, error) {
	if id == "" {
		return nil, fmt.Errorf("Client.PutItem: item ID is required")
	}

	apiURL := c.APIurl
	apiURL.Path += basePath + "/" + id
	if err := p.ParseParams(&apiURL); err != nil {
		return nil, err
	}
	return c.Put(&apiURL)
}

type Poster interface {
	PostItem() ([]byte, error)
	ParseParams(*url.URL) error
}

func (c Client) PostItem(p Poster, basePath string) ([]byte, error) {
	apiURL := c.APIurl
	apiURL.Path += basePath
	if err := p.ParseParams(&apiURL); err != nil {
		return nil, err
	}
	return c.Post(&apiURL)
}

type Getter interface {
	GetAll() ([]byte, error)
	ParseParams(*url.URL) error
}

func (c Client) GetAll(g Getter, basePath string) ([]byte, error) {
	apiURL := c.APIurl
	apiURL.Path += basePath
	if err := g.ParseParams(&apiURL); err != nil {
		return nil, err
	}
	return c.Get(&apiURL)
}
