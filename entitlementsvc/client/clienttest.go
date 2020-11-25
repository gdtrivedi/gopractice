package client

import (
	"fmt"

	"github.com/gdtrivedi/gopractice/fileutil"
)

func EntitlementClientTest() {
	// Read property file for secrets.
	props, e := fileutil.ReadPropertiesFile("/Users/gautam.trivedi/Documents/Work/Projects/OuterEdge/env.properties")
	if e != nil {
		fmt.Println("Error reading properties", e)
		return
	}
	// Get install details by install_name, add a new domain record and update it.
	getEntitlementClient(props)
}

func getEntitlementClient(props fileutil.AppConfigProperties) {
	url := props["auth.prod.url"]
	username := props["auth.prod.username"]
	password := props["auth.prod.password"]
	ProvideEntitlementClient(url, username, password)
}
