package reviewapp

import "time"

type App struct {
	TeamOwner string `json:"TeamOwner"`
	Platform  string `json:"Platform"`
	Name      string `json:"Name"`
	Pool      string `json:"Pool"`
}

type AppInfoResponse struct {
	Platform string
	Pool     string
	Name     string
}

type EnvVar struct {
	Name   string
	Value  string
	Public bool
}

type ReviewAppConfig struct {
	BaseApp string
	Pool    string
	EnvVars []string
}

type ResultSet struct {
	Data    map[string]interface{}
	Timeout time.Duration
	err     error
}
