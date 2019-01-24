package main

import (
	"crypto/tls"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/tsuru/tsuru/cmd"
)

const appName = "test-app-01"

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "tsuru-reviewapp.yml", "path of config file")

	config, err := ParseConfig(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client := cmd.NewClient(httpClient, &cmd.Context{}, &cmd.Manager{})

	u, err := cmd.GetURL(fmt.Sprintf("/apps/%s", config.BaseApp))
	if err != nil {
		fmt.Printf("unable to get URL from target: %v\n", err)
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		fmt.Printf("unable to prepare request: %v\n", err)
		os.Exit(1)
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("unable to fetch app info: %v\n", err)
		os.Exit(1)
	}

	var appInfo AppInfo
	err = json.NewDecoder(resp.Body).Decode(&appInfo)
	if err != nil {
		fmt.Printf("unable to parse app info: %v\n", err)
		os.Exit(1)
	}
	resp.Body.Close()

	u, err = cmd.GetURL(fmt.Sprintf("/apps/%s/env", config.BaseApp))
	if err != nil {
		fmt.Printf("unable to get URL from target: %v\n", err)
		os.Exit(1)
	}

	req, err = http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		fmt.Printf("unable to prepare request: %v\n", err)
		os.Exit(1)
	}

	resp, err = client.Do(req)
	if err != nil {
		fmt.Printf("unable to fetch app env vars: %v\n", err)
		os.Exit(1)
	}

	var baseEnvVars []EnvVar
	err = json.NewDecoder(resp.Body).Decode(&baseEnvVars)
	if err != nil {
		fmt.Printf("unable to parse app info: %v\n", err)
		os.Exit(1)
	}
	resp.Body.Close()

	baseEnv := make(Env)
	for _, envVar := range baseEnvVars {
		baseEnv[envVar.Name] = EnvVar{
			Name:   envVar.Name,
			Value:  envVar.Value,
			Public: envVar.Public,
		}
	}

	env, errs := MergeEnvs(baseEnv, config.Env)
	if errs != nil {
		for _, err := range errs {
			fmt.Printf("unable to collect env vars from config: %v\n", err)
		}
		os.Exit(1)
	}
	fmt.Printf("%+v\n", env)

	// u, err = cmd.GetURL(fmt.Sprintf("/services/instances?app=%s", config.BaseApp))
	// if err != nil {
	// 	fmt.Printf("unable to get URL from target: %v\n", err)
	// 	os.Exit(1)
	// }

	// fmt.Println(u)

	// req, err = http.NewRequest(http.MethodGet, u, nil)
	// if err != nil {
	// 	fmt.Printf("unable to prepare request: %v\n", err)
	// 	os.Exit(1)
	// }

	// resp, err = client.Do(req)
	// if err != nil {
	// 	fmt.Printf("unable to fetch service instances: %v\n", err)
	// 	os.Exit(1)
	// }

	// var services []Service
	// err = json.NewDecoder(resp.Body).Decode(&services)
	// if err != nil {
	// 	fmt.Printf("unable to parse service list: %v\n", err)
	// 	os.Exit(1)
	// }

	// services = FilterServices(services)
	// fmt.Println(services)

	// u, err = cmd.GetURL("/apps")
	// if err != nil {
	// 	fmt.Printf("unable to get URL from target: %v\n", err)
	// 	os.Exit(1)
	// }

	// data := struct {
	// 	Name        string `json:"name"`
	// 	Platform    string `json:"platform"`
	// 	Plan        string `json:"plan"`
	// 	TeamOwner   string `json:"teamOwner"`
	// 	Pool        string `json:"pool"`
	// 	Router      string `json:"router"`
	// 	Description string `json:"description"`
	// }{
	// 	Name:        appName,
	// 	Platform:    appInfo.Platform,
	// 	Plan:        "small",
	// 	TeamOwner:   appInfo.TeamOwner,
	// 	Pool:        appInfo.Pool,
	// 	Router:      appInfo.Router,
	// 	Description: appInfo.Description,
	// }

	// fmt.Printf("%+v\n", data)

	// b := &bytes.Buffer{}
	// err = json.NewEncoder(b).Encode(data)
	// if err != nil {
	// 	fmt.Printf("unable to build app create payload: %v\n", err)
	// 	os.Exit(1)
	// }

	// req, err = http.NewRequest(http.MethodPost, u, b)
	// if err != nil {
	// 	fmt.Printf("unable to prepare request: %v\n", err)
	// 	os.Exit(1)
	// }
	// req.Header.Set("Content-Type", "application/json")

	// fmt.Println("Creating review app...")

	// _, err = client.Do(req)
	// if err != nil {
	// 	fmt.Printf("unable to create app: %v\n", err)
	// 	os.Exit(1)
	// }

	// u, err = cmd.GetURL(fmt.Sprintf("/apps/%s/env", appName))
	// if err != nil {
	// 	fmt.Printf("unable to get URL from target: %v\n", err)
	// 	os.Exit(1)
	// }

	// var envVars []EnvVar
	// for _, envVar := range env {
	// 	envVars = append(envVars, envVar)
	// }

	// envPayload := struct {
	// 	Envs      []EnvVar
	// 	NoRestart bool
	// 	Private   bool
	// }{
	// 	Envs:      envVars,
	// 	NoRestart: true,
	// 	Private:   false,
	// }

	// buf := &bytes.Buffer{}
	// err = json.NewEncoder(buf).Encode(envPayload)
	// if err != nil {
	// 	fmt.Printf("unable to build env vars payload: %v\n", err)
	// 	os.Exit(1)
	// }

	// req, err = http.NewRequest(http.MethodPost, u, buf)
	// if err != nil {
	// 	fmt.Printf("unable to prepare request: %v\n", err)
	// 	os.Exit(1)
	// }
	// req.Header.Set("Content-Type", "application/json")

	// fmt.Println("Settings env vars...")

	// _, err = client.Do(req)
	// if err != nil {
	// 	fmt.Printf("unable to set env vars: %v\n", err)
	// 	os.Exit(1)
	// }
}
