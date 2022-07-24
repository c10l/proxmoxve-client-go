package api

import (
	"encoding/json"

	"github.com/c10l/proxmoxve-client-go/api"
)

const basePath = "/version"

type GetRequest struct {
	Client *api.Client
}

type GetResponseConsole string

const (
	GetResponseConsoleApplet  GetResponseConsole = "applet"
	GetResponseConsoleVV      GetResponseConsole = "vv"
	GetResponseConsoleHTML5   GetResponseConsole = "html5"
	GetResponseConsoleXtermJS GetResponseConsole = "xtermjs"
)

type GetResponse struct {
	Release string `json:"release"`
	RepoID  string `json:"repoid"`
	Version string `json:"version"`
	Console string `json:"console,omitempty"`
}

func (g GetRequest) Do() (*GetResponse, error) {
	var v GetResponse
	apiURL := *g.Client.APIurl
	apiURL.Path += basePath
	resp, err := g.Client.Get(&apiURL)
	if err != nil {
		return nil, err
	}
	return &v, json.Unmarshal(resp, &v)
}
