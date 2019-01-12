package test

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
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
	retCode := m.Run()
	Before()
	os.Exit(retCode)
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
func TestCreateAppReview(t *testing.T) {

	createAppCommand := reviewapp.CreateAppCommand{}
	resultSet := createAppCommand.Run(client)
	builder := strings.Builder{}
	builder.WriteString("review-app")
	builder.WriteString(".gcloud.globoi.com")

	for _, value := range resultSet.Data {
		createCommand := value.(reviewapp.CreateAppCommand)
		assert.Equal(t, "success", createCommand.Status)
		assert.Equal(t, "review-app.gcloud.globoi.com", createCommand.IP, "they should be equal")
	}
}
