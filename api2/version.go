package api2

import (
	"io"
	"strings"
)

const versionBasePath = "/version"

func (c *Client) RetrieveVersion() (io.Reader, error) {
	apiURL := *c.ApiURL
	apiURL.Path += versionBasePath
	resp, err := doGet(c, &apiURL)
	data := strings.NewReader(string(resp))
	return data, err
}
