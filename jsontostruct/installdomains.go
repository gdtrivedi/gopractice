package jsontostruct

import (
	"encoding/json"
	"fmt"
)

// Install represents an install in the DMS service
type Install struct {
	URL     string          `json:"url"`
	Name    string          `json:"name"`
	ID      int             `json:"wpe_id"`
	Activ   bool            `json:"active"`
	Domains []InstallDomain `json:"Domains"`
}

// InstallDomain represents the domain model received from install endpoints
type InstallDomain struct {
	Name      string   `json:"name"`
	Dev       bool     `json:"dev"`
	Redirects []string `json:"redirects"`
}

func InstallDomainsTest() {
	var jsonStr = `
			{
				"url": "https://dms.wpengine.io/dms/v1/installs/gautamtrivedi",
				"wpe_id": 1503045,
			  "active": true,
			  "name": "gautamtrivedi",
			  "domains": [
					{
						"name": "blog.gtrivedi.xyz",
						"dev": false,
						"redirects": []
					},
					{
						"name": "gautamtrivedi.wpengine.com",
						"dev": false,
						"redirects": []
					},
					{
						"name": "gtrivedi.xyz",
						"dev": false,
						"redirects": []
					},
					{
						"name": "test.gtrivedi.xyz",
						"dev": false,
						"redirects": []
					},
					{
						"name": "www.gtrivedi.xyz",
						"dev": false,
						"redirects": []
					}
				]
			}
			`
	install := Install{}
	err := json.Unmarshal([]byte(jsonStr), &install)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Install URL: ", install.URL)
		fmt.Println("Install ID: ", install.ID)
		fmt.Println("install.Domains[0].Name: ", install.Domains[0].Name)
	}
}
