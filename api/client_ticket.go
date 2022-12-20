package api

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type UserPass struct {
	User     string
	Pass     string
	TOTPSeed string
	Ticket   Ticket
}

type Ticket struct {
	CSRFPreventionToken string `json:"CSRFPreventionToken"`
	Ticket              string `json:"ticket"`
	Username            string `json:"username"`
	NeedTFA             int    `json:"NeedTFA"`
}

func NewTicketClient(baseURL, user, pass, tfaSeed string, tlsInsecure bool) (*Client, error) {
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
	ticketBody, err := io.ReadAll(ticketResp.Body)
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
		APIurl: apiURL,
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
