package api2

import (
	"encoding/json"
)

const versionBasePath = "/version"

type Version struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
}

func (c *Client) RetrieveVersion() (*Version, error) {
	apiURL := *c.ApiURL
	apiURL.Path += versionBasePath
	resp, err := doGet(c, &apiURL)
	if err != nil {
		return nil, err
	}
	data := new(Version)
	err = json.Unmarshal(resp, data)
	return data, err
}
