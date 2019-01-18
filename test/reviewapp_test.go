package test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"strings"
	"testing"
	reviewapp "tsuru-reviewapp"

	"github.com/tsuru/tsuru/cmd"

	"github.com/stretchr/testify/assert"
)

var (
	client           *cmd.Client
	dropAppCommand   reviewapp.DropAppCommand
	createAppCommand reviewapp.CreateAppCommand
)

func TestMain(m *testing.M) {
	Init()
	m.Run()
	Before()
}
func Init() {
	fmt.Println("---- Init ----")
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client = cmd.NewClient(httpClient, &cmd.Context{}, &cmd.Manager{})
}

func Before() {
	fmt.Println("---- Before ----")
	//dropAppCommand := reviewapp.DropAppCommand{}
	//dropAppCommand.Run(client)
}

// func TestCreateAppReview(t *testing.T) {
// 	fmt.Println("---- TestCreateAppReview -------")

// 	createAppCommand := reviewapp.CreateAppCommand{}
// 	resultSet := createAppCommand.Run(client, reviewapp.ConfigTsuruTest())
// 	var pathURL = fmt.Sprintf("%s.gcloud.globoi.com", reviewapp.ConfigTsuruTest().BaseApp)

// 	for _, value := range resultSet.Data {
// 		createCommand := value.(reviewapp.CreateAppCommand)
// 		assert.Equal(t, "success", createCommand.Status)
// 		assert.Equal(t, pathURL, createCommand.IP, "they should be equal")
// 	}
// }

/// dbaas services criam bases e o nome permanece no sistema verificar uma possibilidade de gerar um nomenclatura que
// seria facil de apagar  pela equipe posteriormente
func TestCommandServiceAdd(t *testing.T) {
	fmt.Println("---- TestCommandServiceAdd -------")

	serviceAddCommand := reviewapp.AddServiceAppCommand{}
	resultSet := serviceAddCommand.Run(client, reviewapp.ConfigTsuruTest())

	for _, value := range resultSet.Data {

		statusCode := value.(int)
		fmt.Println(statusCode)
		assert.Equal(t, http.StatusCreated, statusCode)
	}
}

func TestCommandDropApp(t *testing.T) {
	fmt.Println("---- TestCommandDropApp -------")

	dropAppCommand := reviewapp.DropAppCommand{}
	resultSet := dropAppCommand.Run(client, reviewapp.ConfigTsuruTest())

	for _, value := range resultSet.Data {
		dropAppCommand := value.(reviewapp.DropAppCommand)
		assert.True(t, strings.Contains(dropAppCommand.Message, "Removing application"))
	}
}
