// Test global ingress -> cloudflare API calls.
package gi

import (
	"fmt"

	"github.com/cloudflare/cloudflare-go"
	"github.com/gdtrivedi/gopractice/fileutil"
)

func GICFTest() {
	env := "stage"
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
	customHostnameByName(api, zoneId)
}

func customHostnameById(api *cloudflare.API, zoneId string) {
	id := "004468a5-25da-4113-b4a5-afea993807db"
	hostname, err := api.CustomHostname(zoneId, id)
	if err != nil {
		fmt.Println("CustomHostname Error >>> ", err)
	} else {
		fmt.Println("hostname.Hostname >>> ", hostname.Hostname)
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
