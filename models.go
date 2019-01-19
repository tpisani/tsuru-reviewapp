package reviewapp

import "time"

type App struct {
	TeamOwner string `json:"TeamOwner"`
	Platform  string `json:"Platform"`
	Name      string `json:"Name"`
	Pool      string `json:"Pool"`
}

type Service struct {
	PlanName    string `json:"plan"`
	Owner       string `json:"owner"`
	Description string `json:"Description"`
	Name        string `json:"Name"`
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

	Dbaas DbaasMySQL `yaml:"dbaas-mysql"`
}
type DbaasMySQL struct {
	Service     string `yaml:"service"`
	ServiceName string `yaml:"service-name"`
	PlanName    string `yaml:"plan-name"`
	TeamOwner   string `yaml:"team-owner"`
}
type ResultSet struct {
	Data    map[string]interface{}
	Timeout time.Duration
	err     error
}
