// Test global ingress -> cloudflare API calls.
package gi

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gdtrivedi/gopractice/fileutil"
)

func GICFTest() {
	env := "stage"
	//env := "prod"
	// Read property file for secrets.
	props, e := fileutil.ReadPropertiesFile("/Users/gautam.trivedi/Documents/Work/Projects/OuterEdge/env.properties")
	if e != nil {
		fmt.Println("Error reading properties", e)
		return
	}
	cfKey := props[env+".gi.cf.key"]
	cfEmail := props[env+".gi.cf.email"]
	zoneId := props[env+".gi.cf.zoneID"]
	api, err := cloudflare.New(cfKey, cfEmail)

	if err != nil {
		fmt.Println("Error >>> ", err)
	} else {
		restAPICaller(api, zoneId)
	}

	// API instance using token.
	//api2, err2 := cloudflare.NewWithAPIToken(config.APIToken, cloudflare.HTTPClient(&http.Client{}))
}

func restAPICaller(api *cloudflare.API, zoneId string) {
	customHostnameById(api, zoneId)
	//customHostnameByName(api, zoneId)
	//createCustomHostname(api, zoneId)
}

func customHostnameById(api *cloudflare.API, zoneId string) {
	id := "e7e00cca-4041-45f1-90ed-a427cbf44275"
	hostname, err := api.CustomHostname(zoneId, id)
	if err != nil {
		fmt.Println("CustomHostname Error >>> ", err)
	} else {
		fmt.Println("hostname>>> ", hostname.SSL)
		//fmt.Println("hostname.Hostname >>> ", hostname.Hostname)
	}
}

func customHostnameByName(api *cloudflare.API, zoneId string) {
	hostname := "stagingfarmingresspagereadtest019.wptestinstall.com"
	customHostnames, resultInfo, err := api.CustomHostnames(zoneId, 1, cloudflare.CustomHostname{Hostname: hostname})
	if err != nil {
		fmt.Println("CustomHostnames Error >>> ", err)
	} else {
		fmt.Println("Total: ", resultInfo.Count)
		for _, customHostname := range customHostnames {
			fmt.Println("customHostname.ID: ", customHostname.ID)
		}
	}
}
func createCustomHostname(api *cloudflare.API, zoneId string) {
	hostname := cloudflare.CustomHostname{
		Hostname:           "www.gautam03.xyz",
		CustomOriginServer: "farmingresssli.wpenginepoweredstaging.com",
		SSL: cloudflare.CustomHostnameSSL{
			Method: "http",
			Type:   "dv",
			Settings: cloudflare.CustomHostnameSSLSettings{
				MinTLSVersion: "1.2",
			},
		},
	}
	chResp, err := api.CreateCustomHostname(zoneId, hostname)
	if err != nil {
		fmt.Println("CustomHostnames Error >>> ", err)
	} else {
		fmt.Println("chResp.Result: ", chResp.Result)
		fmt.Println("chResp.Result: ", chResp.Result.SSL)
	}
	/*
		/private/var/folders/lz/f2bvlq_925nfcdv50vkzgm_40000gp/T/___go_build_github_com_gdtrivedi_gopractice
		chResp.Result:  {e7e00cca-4041-45f1-90ed-a427cbf44275 www.gautam03.xyz farmingresssli.wpenginepoweredstaging.com {initializing http dv   false   digicert {  1.2 []} []  } map[] pending [] {txt _cf-custom-hostname.www.gautam03.xyz b97facb8-7adf-41c0-ab78-712367ae90e4} {http://www.gautam03.xyz/.well-known/cf-custom-hostname-challenge/e7e00cca-4041-45f1-90ed-a427cbf44275 b97facb8-7adf-41c0-ab78-712367ae90e4} 2020-12-18 20:41:04.195091 +0000 UTC}
		chResp.Result:  {initializing http dv   false   digicert {  1.2 []} []  }

		Process finished with exit code 0

	*/
}
