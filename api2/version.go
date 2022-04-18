package api2

type Version struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
	Console string `json:"console,omitempty"`
}

const versionBasePath = "/version"

func (c *Client) RetrieveVersion() (*Version, error) {
	data := new(Version)
	apiURL := *c.ApiURL
	apiURL.Path += versionBasePath
	err := doGet(c, data, &apiURL)
	return data, err
}
