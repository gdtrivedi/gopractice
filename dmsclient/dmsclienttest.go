package dmsclient

import (
	"fmt"
	"net/http"

	"github.com/gdtrivedi/gopractice/fileutil"
	"github.com/gdtrivedi/gopractice/httputil"
)

func DMSClientTest() {
	// Read property file for secrets.
	props, e := fileutil.ReadPropertiesFile("/Users/gautam.trivedi/Documents/Work/Projects/OuterEdge/env.properties")
	if e != nil {
		fmt.Println("Error reading properties", e)
		return
	}
	// Get install details by install_name, add a new domain record and update it.
	getByInstallNameAndAddDomainRecord(props)
}

func getByInstallNameAndAddDomainRecord(props fileutil.AppConfigProperties) {
	httpClient := httputil.SetClientDefaults(&http.Client{})
	httpu := httputil.NewUtility(httpClient)
	client := NewClient(httpClient, props["dms.dev.authKey"], props["dms.dev.baseURL"], httpu)
	install, err := client.FetchInstallInfoByInstallName("gientprod")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Install ID:", install.ID)
		// Append <install-name>.<wpenginepowered.com> domain record.
		install.Domains = append(install.Domains, InstallDomain{
			Name: "testinstall" + "." + "testdomain.com",
			Dev:  false,
		})
		//call DMS to update install. This call will eventually add domain.
		flag, err := client.UpdateInstallInfo(install)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Install Updated:", flag)
		}
	}
}
