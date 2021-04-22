package dmsclient

import (
	"fmt"
	"net/http"

	"github.com/golang/protobuf/proto"

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
	//getByInstallNameAndAddDomainRecord(props)
	//getDomainInfoByID(props)
	//getAtlasEnvInfoByID(props)
	getDomainsByName(props)
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

func getDomainInfoByID(props fileutil.AppConfigProperties) {
	httpClient := httputil.SetClientDefaults(&http.Client{})
	httpu := httputil.NewUtility(httpClient)
	client := NewClient(httpClient, props["dms.stage.authKey"], props["dms.stage.baseURL"], httpu)
	domainInfoByID, err := client.GetDomainInfoByID(proto.Int32(1249636))
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("domainInfoByID.ID: ", *(domainInfoByID.ID))
		fmt.Println("domainInfoByID.Name: ", domainInfoByID.Name)
	}
}

func getAtlasEnvInfoByID(props fileutil.AppConfigProperties) {
	httpClient := httputil.SetClientDefaults(&http.Client{})
	httpu := httputil.NewUtility(httpClient)
	client := NewClient(httpClient, props["dms.stage.authKey"], props["dms.stage.baseURL"], httpu)
	envInfo, err := client.FetchAtlasEnvInfoByID("7a81slxd19")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("envInfo.ID: ", envInfo.ID)
		fmt.Println("envInfo.Name: ", envInfo.OriginServer)
		fmt.Println("len(envInfo.Domains): ", len(envInfo.Domains))
	}
}

func getDomainsByName(props fileutil.AppConfigProperties) {
	httpClient := httputil.SetClientDefaults(&http.Client{})
	httpu := httputil.NewUtility(httpClient)
	client := NewClient(httpClient, props["dms.dev.authKey"], props["dms.dev.baseURL"], httpu)
	domains, err := client.GetDomainsByName("a0clone200003.wpengine.com")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		for _, domain := range domains {
			fmt.Println("domain.ID: ", domain.ID)
			fmt.Println("domain.Name: ", domain.Name)
		}
	}
}
