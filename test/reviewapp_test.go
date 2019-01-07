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

var client *cmd.Client

func TestMain(m *testing.M) {
	Init()
	retCode := m.Run()
	Before()
	os.Exit(retCode)
}
func Init() {
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}
	client = cmd.NewClient(httpClient, &cmd.Context{}, &cmd.Manager{})
}

func Before() {

}
func TestCreateAppReview(t *testing.T) {

	createAppCommand := reviewapp.CreateAppCommand{}
	dropAppCommand := reviewapp.DropAppCommand{}

	resultSet := createAppCommand.Run(client)
	builder := strings.Builder{}

	for _, value := range resultSet.Data {
		createCommand := value.(reviewapp.CreateAppCommand)
		builder.WriteString(reviewapp.ConfigTsuru().BaseApp)
		builder.WriteString(".gcloud.globoi.com")
		fmt.Println(builder.String())
		assert.Equal(t, builder.String(), createCommand.IP, "they should be equal")
		assert.Equal(t, "success", createAppCommand.Status)
	}
	dropAppCommand.Run(client)
}
