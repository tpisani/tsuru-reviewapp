package reviewapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/tsuru/tsuru/cmd"
)

type AddServiceAppCommand struct {
}

func (p *AddServiceAppCommand) Run(client *cmd.Client, review ReviewAppConfig) ResultSet {
	urlPath, err := cmd.GetURL(fmt.Sprintf("/services/%s/instances", review.Dbaas.Service))
	if err != nil {
		fmt.Println("AddServiceAppCommand to get URL from target")
		os.Exit(1)
	}

	fmt.Println(review.Dbaas.ServiceName)

	data := Service{}
	data.Name = review.Dbaas.ServiceName
	data.Description = "banco mysql para teste"
	data.Owner = review.Dbaas.TeamOwner
	data.PlanName = review.Dbaas.PlanName

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

	dataResponse := map[string]interface{}{
		"status": resp.StatusCode,
	}
	resultSet := ResultSet{
		Data: dataResponse,
	}

	return resultSet
}

type RemoveServiceAppCommand struct {
}

func (p *RemoveServiceAppCommand) Run(client *cmd.Client, review ReviewAppConfig) ResultSet {

	urlPath, err := cmd.GetURL(fmt.Sprintf("/services/%s/instances/%s", review.Dbaas.Service, review.Dbaas.ServiceName))

	if err != nil {
		fmt.Println("unable RemoveServiceAppCommand to get URL from target")
		os.Exit(1)
	}
	req, _ = http.NewRequest(http.MethodDelete, urlPath, nil)
	resp, err = client.Do(req)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	dataResponse := map[string]interface{}{
		"status": resp.StatusCode,
	}
	resultSet := ResultSet{
		Data: dataResponse,
	}

	return resultSet
}
