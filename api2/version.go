package api2

type Version struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
	Console string `json:"console,omitempty"`
}

const versionBasePath = "/version"

func (c *Client) RetrieveVersion() (*Version, error) {
	var data *Version
	apiURL := *c.ApiURL
	apiURL.Path += versionBasePath
	return doGet(c, data, &apiURL)
}
