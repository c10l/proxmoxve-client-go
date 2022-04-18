package api2

type Version struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
	Console string `json:"console,omitempty"`
}

const versionBasePath = "/version"

func (c *Client) GetVersion() (*Version, error) {
	var data *Version
	// url := fmt.Sprintf(c.BaseURL + versionBasePath)
	apiURL := *c.ApiURL
	apiURL.Path += versionBasePath

	return doGet(c, data, &apiURL)
}
