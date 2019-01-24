package main

type AppInfo struct {
	Name        string
	TeamOwner   string
	Platform    string
	Pool        string
	Router      string
	Description string
}

type AppInfoFetcher interface {
	FetchAppInfo(name string) (*AppInfo, error)
}
