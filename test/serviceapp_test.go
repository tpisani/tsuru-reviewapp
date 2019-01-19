package test

import (
	"fmt"
	"net/http"
	"testing"
	reviewapp "tsuru-reviewapp"

	"github.com/stretchr/testify/assert"
)

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

func TestCommandServiceRemove(t *testing.T) {
	fmt.Println("---- TestCommandServiceRemove -------")

	serviceRemoveCommand := reviewapp.RemoveServiceAppCommand{}
	resultSet := serviceRemoveCommand.Run(client, reviewapp.ConfigTsuruTest())

	for _, value := range resultSet.Data {

		statusCode := value.(int)
		fmt.Println(statusCode)
		assert.Equal(t, http.StatusOK, statusCode)
	}
}
