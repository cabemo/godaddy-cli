// Package config sets up all the configuration for GoDaddy CLI
package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os/user"

	"github.com/oze4/godaddygo"
)

// APICredentials used in each request
type APICredentials struct {
	Key    string `json:"key"`
	Secret string `json:"secret"`
}

// Credentials reads the credentials files located in $HOME/.config/godaddy/credentials.json and returns a Credentials
func Credentials() (APICredentials, error) {
	var creds APICredentials
	// Get current user in order to get credentials
	user, err := user.Current()

	if err != nil {
		panic(err.Error())
	}

	data, err := ioutil.ReadFile(user.HomeDir + "/.config/godaddy/credentials.json")

	if err != nil {
		fmt.Println("Error while reading config file ($HOME/.config/godaddy/credentials.json)")
		panic(nil)
	}

	err = json.Unmarshal(data, &creds)

	if err != nil {
		return creds, err
	}

	return creds, nil
}

// GoDaddy creates a new oze4/godaddgo.API instance with the configuration applied
func GoDaddy() (godaddygo.V1, error) {
	creds, err := Credentials()
	if err != nil {
		return nil, err
	}

	api, err := godaddygo.NewProduction(creds.Key, creds.Secret)
	if err != nil {
		return nil, err
	}
	godaddy := api.V1()

	return godaddy, nil
}
