package main

type Settings struct {
	APIKey       string `json:"api_key"`
	Organization string `json:"organization"`
	Repositories []struct {
		Name string `json:"name"`
	} `json:"repositories"`
}

var settingsJson = `{
	"api_key": "your_key",
	"organization": "petterr-dot",
	"repositories": [
	  {
		"name": "gofetchpr"
	  },
	  {
		"name": "somerepo"
	  }
	]
  }`
