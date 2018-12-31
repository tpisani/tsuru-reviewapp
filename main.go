package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/tsuru/tsuru/cmd"
	yaml "gopkg.in/yaml.v2"
)

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

var (
	url  string
	err  error
	req  *http.Request
	resp *http.Response
)

func filterEnvVars(envVars []EnvVar, names ...string) []EnvVar {
	filtered := make([]EnvVar, 0)

	for _, name := range names {
		for _, v := range envVars {
			if v.Public && v.Name == name {
				filtered = append(filtered, EnvVar{
					Name:   v.Name,
					Value:  v.Value,
					Public: v.Public,
				})
			}
		}
	}

	return filtered
}

func configTsuru() ReviewAppConfig {

	token := os.Getenv("TSURU_TOKEN")
	if token == "" {
		fmt.Println("missing Tsuru token")
		os.Exit(1)
	}

	target := os.Getenv("TSURU_TARGET")
	if target == "" {
		fmt.Println("missing Tsuru target")
		os.Exit(1)
	}

	fmt.Println(token, target)

	f, err := os.Open("./tsuru-reviewapp.yml")
	if err != nil {
		fmt.Println("no tsuru-reviewapp.yml")
		os.Exit(1)
	}

	var config ReviewAppConfig
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		fmt.Println("unable to parse config file")
		os.Exit(1)
	}
	f.Close()

	return config
}

func main() {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	clientConfig := ClientConfig{
		Client:  *httpClient,
		Context: cmd.Context{},
		Manager: cmd.Manager{},
	}
	execCommands(clientConfig)
}
