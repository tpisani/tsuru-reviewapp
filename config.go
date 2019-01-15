package reviewapp

import (
	"fmt"
	"net/http"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var (
	urlpath string
	err     error
	req     *http.Request
	resp    *http.Response
)

func FilterEnvVars(envVars []EnvVar, names ...string) []EnvVar {
	filtered := make([]EnvVar, 0)

	for _, name := range names {
		for _, v := range envVars {
			if v.Public && v.Name == name {
				filtered = append(filtered, EnvVar{
					Name:   v.Name,
					Value:  v.Value,
					Public: v.Public,
				})
			}
		}
	}
	return filtered
}

func ConfigTsuru() ReviewAppConfig {

	token := os.Getenv("TSURU_TOKEN")
	if token == "" {
		fmt.Println("missing Tsuru token")
		os.Exit(1)
	}

	target := os.Getenv("TSURU_TARGET")
	if target == "" {
		fmt.Println("missing Tsuru target")
		os.Exit(1)
	}

	f, err := os.Open("./tsuru-reviewapp.yml")
	if err != nil {
		fmt.Println("no tsuru-reviewapp.yml")
		os.Exit(1)
	}

	var config ReviewAppConfig
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		fmt.Println("unable to parse config file")
		os.Exit(1)
	}
	f.Close()

	return config
}

func ConfigTsuruTest() ReviewAppConfig {

	token := os.Getenv("TSURU_TOKEN")
	if token == "" {
		fmt.Println("missing Tsuru token")
		os.Exit(1)
	}

	target := os.Getenv("TSURU_TARGET")
	if target == "" {
		fmt.Println("missing Tsuru target")
		os.Exit(1)
	}

	f, err := os.Open("../tsuru-reviewapp.yml")
	if err != nil {
		fmt.Println("no tsuru-reviewapp.yml")
		os.Exit(1)
	}

	var config ReviewAppConfig
	err = yaml.NewDecoder(f).Decode(&config)
	if err != nil {
		fmt.Println("unable to parse config file")
		os.Exit(1)
	}
	f.Close()

	return config
}
