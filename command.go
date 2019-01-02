package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/tsuru/tsuru/cmd"
)

type ResultSet struct {
	Data    map[string]interface{}
	Timeout time.Duration
	err     error
}

type Command interface {
	Run(newClient *cmd.Client) ResultSet
	//RoolBack() string
}

type AppInfoCommand struct{}

func (p *AppInfoCommand) Run(client *cmd.Client) ResultSet {

	url, err := cmd.GetURL(fmt.Sprintf("/apps/%s", configTsuru().BaseApp))
	if err != nil {
		fmt.Println("unable to get URL from target")
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
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

	appInfo := AppInfoResponse{}

	err = json.NewDecoder(resp.Body).Decode(&appInfo)
	if err != nil {
		fmt.Printf("unable to parse app info: %v\n", err)
		os.Exit(1)
	}

	result := map[string]interface{}{
		"name":     appInfo.Name,
		"platform": appInfo.Platform,
		"pool":     appInfo.Pool,
	}

	resultSet := ResultSet{
		Data: result,
	}

	return resultSet
}

type GetEnvCommand struct{}

func (p *GetEnvCommand) Run(client *cmd.Client) ResultSet {

	url, err = cmd.GetURL(fmt.Sprintf("/apps/%s/env", configTsuru().BaseApp))

	if err != nil {
		fmt.Println("unable to get URL from target")
		os.Exit(1)
	}

	req, _ = http.NewRequest(http.MethodGet, url, nil)
	resp, err = client.Do(req)
	fmt.Println(resp.Body)

	if err != nil {
		fmt.Println("Contain a non-nil Body")
		os.Exit(1)
	}

	defer resp.Body.Close()

	var envVars []EnvVar

	json.NewDecoder(resp.Body).Decode(&envVars)

	envVars = filterEnvVars(envVars, "NODE_ENV", "FEATURES")
	envs := map[string]interface{}{
		"envs": envVars,
	}
	resultSet := ResultSet{
		Data: envs,
	}

	return resultSet
}

/*
type CreateCommand struct{}

func (p *CreateCommand) Run() string {
	return "CreateCommand"
}
func (p *CreateCommand) RoolBack() string {
	return "RoolBack CreateCommand"
}

type BindCommand struct{}

func (p *BindCommand) Run() string {
	return "BindCommand"
}
func (p *BindCommand) RoolBack() string {
	return "RoolBack BindCommand"
}

type UnBindCommand struct{}

func (p *UnBindCommand) Run() string {
	return "UnBindCommand"
}
func (p *UnBindCommand) RoolBack() string {
	return "RoolBack UnBindCommand"
}

type DropCommand struct{}

func (p *DropCommand) Run() string {
	return "DropCommand"
}

func (p *DropCommand) RoolBack() string {
	return "RoolBack DropCommand"
}
*/
func execByName(name string, client *cmd.Client) {
	commands := map[string]Command{
		"info": &AppInfoCommand{},
		"env":  &GetEnvCommand{},
	}
	if command := commands[name]; command == nil {
		fmt.Println("No such command found, throw error?")
	} else {
		command.Run(client)
	}
}

func execCommands(client *cmd.Client) {
	// Register commands
	commands := [...]Command{
		&AppInfoCommand{},
		&GetEnvCommand{},
	}

	for _, command := range commands {
		fmt.Println(command.Run(client))
	}
}
