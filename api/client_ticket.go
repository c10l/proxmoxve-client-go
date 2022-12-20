package api

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/pquerna/otp/totp"
)

type UserPass struct {
	User     string
	Pass     string
	TOTPSeed string
}

type Ticket struct {
	CSRFPreventionToken string `json:"CSRFPreventionToken"`
	Ticket              string `json:"ticket"`
	Username            string `json:"username"`
	NeedTFA             int    `json:"NeedTFA"`
}

func NewTicketClient(baseURL, user, pass, totpSeed string, tlsInsecure bool) (*Client, error) {
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
	formData := url.Values{
		"username":   {user},
		"password":   {url.QueryEscape(pass)},
		"new-format": {"1"},
	}
	ticket, err := getTicket(apiURL, formData, httpClient)
	if err != nil {
		return nil, err
	}

	// If the root user has TFA enabled, we need to re-call the ticket API with the TOTP
	if ticket.NeedTFA == 1 {
		if totpSeed == "" {
			return nil, errors.New("TFA required - please inform TOTP seed")
		}
		totp, err := getTOTP(totpSeed)
		if err != nil {
			return nil, err
		}

		// Retrieve TOTP ticket
		formData := url.Values{
			"username": {user},
			"password": {"totp:" + totp},
		}
		formData.Set("tfa-challenge", ticket.Ticket)
		ticket, err = getTicket(apiURL, formData, httpClient)
		if err != nil {
			return nil, err
		}
	}

	client := &Client{
		APIurl:      *apiURL,
		Ticket:      ticket,
		TLSInsecure: tlsInsecure,
		HTTPClient:  httpClient,
	}

	return client, nil
}

func getTicket(apiURL *url.URL, formData url.Values, httpClient *http.Client) (*Ticket, error) {
	ticketURL := *apiURL
	ticketURL.Path += "/access/ticket"
	ticketResp, err := httpClient.PostForm(ticketURL.String(), formData)
	if err != nil {
		return nil, err
	}
	defer ticketResp.Body.Close()
	ticketBody, err := io.ReadAll(ticketResp.Body)
	if err != nil {
		return nil, err
	}
	if ticketResp.StatusCode != http.StatusOK {
		fmt.Print("salsifufu")
		return nil, fmt.Errorf("%s: %s", ticketResp.Status, ticketBody)
	}
	ticketRaw, err := parseData(ticketBody)
	if err != nil {
		return nil, err
	}
	ticket := &Ticket{}
	if err := json.Unmarshal(ticketRaw, &ticket); err != nil {
		return nil, err
	}
	return ticket, nil
}

func getTOTP(totpSeed string) (string, error) {
	totp, err := totp.GenerateCode(totpSeed, time.Now())
	return totp, err
}
