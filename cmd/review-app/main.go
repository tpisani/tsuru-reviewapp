package main

import (
	"crypto/tls"
	"net/http"
	reviewapp "tsuru-reviewapp"

	"github.com/tsuru/tsuru/cmd"
)

func main() {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client := cmd.NewClient(httpClient, &cmd.Context{}, &cmd.Manager{})
	reviewapp.ExecCommands(client)
}
