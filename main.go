package main

type AppInfoResponse struct {
	Platform string
	Pool     string
}

type EnvVar struct {
	Name   string
	Value  string
	Public bool
}

type ReviewAppConfig struct {
	BaseApp string
	Pool    string
	EnvVars []string
}

func filterEnvVars(envVars []EnvVar, names ...string) []EnvVar {
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

func main() {
	execCommands()
}

/*
func main() {
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

	fmt.Println(token, target)

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

	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	client := cmd.NewClient(httpClient, &cmd.Context{}, &cmd.Manager{})

	u, err := cmd.GetURL(fmt.Sprintf("/apps/%s", config.BaseApp))
	if err != nil {
		fmt.Println("unable to get URL from target")
		os.Exit(1)
	}

	req, err := http.NewRequest(http.MethodGet, u, nil)
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

	var data AppInfoResponse
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Printf("unable to parse app info: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(data.Platform, data.Pool)

	u, err = cmd.GetURL(fmt.Sprintf("/apps/%s/env", config.BaseApp))
	if err != nil {
		fmt.Println("unable to get URL from target")
		os.Exit(1)
	}

	req, _ = http.NewRequest(http.MethodGet, u, nil)
	resp, _ = client.Do(req)
	defer resp.Body.Close()

	var envVars []EnvVar
	json.NewDecoder(resp.Body).Decode(&envVars)
	envVars = filterEnvVars(envVars, "NODE_ENV", "FEATURES")
	fmt.Println(envVars)
}
*/
