package reviewapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tsuru/tsuru/cmd"
)

type Command interface {
	Run(newClient *cmd.Client, review ReviewAppConfig) ResultSet
}

type AppInfoCommand struct{}

func (p *AppInfoCommand) Run(client *cmd.Client, review ReviewAppConfig) ResultSet {

	urlPath, err := cmd.GetURL(fmt.Sprintf("/apps/%s", review.BaseApp))
	if err != nil {
		fmt.Println("unable to get URL from target")
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
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

func (p *GetEnvCommand) Run(client *cmd.Client, review ReviewAppConfig) ResultSet {

	urlPath, err := cmd.GetURL(fmt.Sprintf("/apps/%s/env", review.BaseApp))

	if err != nil {
		fmt.Println("unable to get URL from target")
		os.Exit(1)
	}

	req, _ = http.NewRequest(http.MethodGet, urlPath, nil)
	resp, err = client.Do(req)
	fmt.Println(resp.Body)

	if err != nil {
		fmt.Println("Contain a non-nil Body")
		os.Exit(1)
	}

	defer resp.Body.Close()

	var envVars []EnvVar

	json.NewDecoder(resp.Body).Decode(&envVars)

	envVars = FilterEnvVars(envVars, "NODE_ENV", "FEATURES")
	envs := map[string]interface{}{
		"envs": envVars,
	}
	resultSet := ResultSet{
		Data: envs,
	}

	return resultSet
}

type CreateAppCommand struct {
	IP            string
	RepositoryURL string
	Status        string
}

func (p *CreateAppCommand) Run(client *cmd.Client, review ReviewAppConfig) ResultSet {

	urlPath, err := cmd.GetURL(fmt.Sprintf("/apps"))
	if err != nil {
		fmt.Println("unabCreateAppCommandle to get URL from target")
		os.Exit(1)
	}

	data := App{}
	data.TeamOwner = "backend_produtos_globosat"
	data.Platform = "python"
	data.Name = "review-app"
	data.Pool = "globosat"

	dataPost, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req, _ = http.NewRequest(http.MethodPost, urlPath, bytes.NewBuffer(dataPost))
	req.Header.Set("Content-Type", "application/json")
	resp, err = client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	appCreate := CreateAppCommand{}

	json.NewDecoder(resp.Body).Decode(&appCreate)

	dataResponse := map[string]interface{}{
		"app-create": appCreate,
	}
	resultSet := ResultSet{
		Data: dataResponse,
	}

	return resultSet
}

type DropAppCommand struct {
	Message string
}

func (p *DropAppCommand) Run(client *cmd.Client, review ReviewAppConfig) ResultSet {
	urlPath, err := cmd.GetURL(fmt.Sprintf("/apps/%s", review.BaseApp))
	if err != nil {
		fmt.Println("unableDropAppCommandle to get URL from target")
		os.Exit(1)
	}
	req, _ = http.NewRequest(http.MethodDelete, urlPath, nil)
	resp, err = client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	dropAppCommand := DropAppCommand{}

	json.NewDecoder(resp.Body).Decode(&dropAppCommand)

	dataResponse := map[string]interface{}{
		"app-drop": dropAppCommand,
	}
	resultSet := ResultSet{
		Data: dataResponse,
	}

	return resultSet
}

func ExecByName(name string, client *cmd.Client) {
	commands := map[string]Command{
		"info": &AppInfoCommand{},
		"env":  &GetEnvCommand{},
		"drop": &DropAppCommand{},
	}
	if command := commands[name]; command == nil {
		fmt.Println("No such command found, throw error?")
	} else {
		command.Run(client, ConfigTsuru())
	}
}

func ExecCommands(client *cmd.Client) {

	commands := [...]Command{
		//&AppInfoCommand{},
		//&GetEnvCommand{},
		//&CreateAppCommand{},
		&AddServiceAppCommand{},
	}

	for _, command := range commands {
		fmt.Println(command.Run(client, ConfigTsuru()))
	}
}
