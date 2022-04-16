package api2

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Version struct {
	Data VersionData
}

type VersionData struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
	Console string `json:"console,omitempty"`
}

const versionBasePath = "/version"

func (c *Client) GetVersion() (*Version, error) {
	var data Version
	url := fmt.Sprintf(c.BaseURL + versionBasePath)

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
	return &data, nil
}
