package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tsuru/tsuru/cmd"
)

type ClientConfig struct {
	Client  http.Client
	Context cmd.Context
	Manager cmd.Manager
}

type Command interface {
	Run(clientConfig *ClientConfig)
	//RoolBack() string
}

type AppInfoCommand struct{}

func (p *AppInfoCommand) Run(clientConfig *ClientConfig) {

	client := cmd.NewClient(&clientConfig.Client, &clientConfig.Context, &clientConfig.Manager)

	u, err := cmd.GetURL(fmt.Sprintf("/apps/%s", configTsuru().BaseApp))
	if err != nil {
		fmt.Println("unable to get URL from target")
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		fmt.Println("unable to prepare request")
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("unable to fetch app info: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("non-200 response from Tsuru")
		os.Exit(1)
	}

	var data AppInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("unable to parse app info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(data.Platform, data.Pool)
}

type GetEnvCommand struct{}

func (p *GetEnvCommand) Run(clientConfig *ClientConfig) {

	url, err = cmd.GetURL(fmt.Sprintf("/apps/%s/env", configTsuru().BaseApp))

	if err != nil {
		fmt.Println("unable to get URL from target")
		os.Exit(1)
	}

	req, _ = http.NewRequest(http.MethodGet, url, nil)
	resp, _ = clientConfig.Client.Do(req)
	defer resp.Body.Close()

	var envVars []EnvVar
	json.NewDecoder(resp.Body).Decode(&envVars)
	envVars = filterEnvVars(envVars, "NODE_ENV", "FEATURES")
	fmt.Println(envVars)
}

/*
type CreateCommand struct{}

func (p *CreateCommand) Run() string {
	return "CreateCommand"
}
func (p *CreateCommand) RoolBack() string {
	return "RoolBack CreateCommand"
}

func (p *CreateCommand) Info() string {
	return "Info CreateCommand"
}

func (p *CreateCommand) Timeout() time.Duration {
	return 12122
}
*/

/*
type BindCommand struct{}

func (p *BindCommand) Run() string {
	return "BindCommand"
}
func (p *BindCommand) RoolBack() string {
	return "RoolBack BindCommand"
}
func (p *BindCommand) Info() string {
	return "Info BindCommand"
}

func (p *BindCommand) Timeout() time.Duration {
	return 12122
}



type UnBindCommand struct{}

func (p *UnBindCommand) Run() string {
	return "UnBindCommand"
}
func (p *UnBindCommand) RoolBack() string {
	return "RoolBack UnBindCommand"
}

func (p *UnBindCommand) Info() string {
	return "Info UnBindCommand"
}

func (p *UnBindCommand) Timeout() time.Duration {
	return 12122
}

type DropCommand struct{}

func (p *DropCommand) Run() string {
	return "DropCommand"
}

func (p *DropCommand) RoolBack() string {
	return "RoolBack DropCommand"
}

func (p *DropCommand) Info() string {
	return "Info DropCommand"
}

func (p *DropCommand) Timeout() time.Duration {
	return 12122
}
*/
func execByName(name string, clientConfig ClientConfig) {
	commands := map[string]Command{
		"info": &AppInfoCommand{},
		"envs": &GetEnvCommand{},
	}
	if command := commands[name]; command == nil {
		fmt.Println("No such command found, throw error?")
	} else {
		command.Run(&clientConfig)
	}
}

func execCommands(clientConfig ClientConfig) {
	// Register commands
	commands := [...]Command{
		&AppInfoCommand{},
	}

	for _, command := range commands {
		command.Run(&clientConfig)
		fmt.Println()
	}
}
