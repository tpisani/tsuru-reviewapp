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
	token   string
	target  string
	config  ReviewAppConfig
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

func configToken() {
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

}
func ConfigTsuru() ReviewAppConfig {
	configToken()
	file, err := os.Open("./tsuru-reviewapp.yml")
	if err != nil {
		fmt.Println("no tsuru-reviewapp.yml")
		os.Exit(1)
	}
	return getReviewAppConfig(config, file)
}

func ConfigTsuruTest() ReviewAppConfig {
	configToken()
	file, err := os.Open("../tsuru-reviewapp.yml")
	if err != nil {
		fmt.Println("no tsuru-reviewapp.yml")
		os.Exit(1)
	}

	return getReviewAppConfig(config, file)
}

func getReviewAppConfig(config ReviewAppConfig, file *os.File) ReviewAppConfig {
	err = yaml.NewDecoder(file).Decode(&config)
	if err != nil {
		fmt.Println("unable to parse config file")
		os.Exit(1)
	}
	file.Close()
	return config
}
