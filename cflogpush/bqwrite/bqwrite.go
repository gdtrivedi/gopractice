package bqwrite

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

// OfferingsFilter represents the available fields to filter all offerings by.
type OfferingsFilter struct {
	InstallName     *string `url:"install_name,omitempty"`
	Limit           *int32  `url:"limit,omitempty"`
	Offset          *int32  `url:"offset,omitempty"`
	Reconciled      *bool   `url:"reconciled,omitempty"`
	SecureSubdomain *string `url:"secure_subdomain,omitempty"`
}

// Offering represents an offering record.
type Offering struct {
	URL              string    `json:"url"`
	UUID             string    `json:"uuid"`
	OfferingTypeUUID string    `json:"offering_type_uuid"`
	InstallName      string    `json:"install_name"`
	AccountID        int       `json:"account_id"`
	SecureSubdomain  string    `json:"secure_subdomain"`
	Removed          bool      `json:"removed"`
	AccountEntitled  bool      `json:"account_entitled"`
	Reconciled       bool      `json:"reconciled"`
	SiteConfigSet    bool      `json:"site_config_set"`
	CreatedOn        time.Time `json:"created_on,omitempty"`
	UpdatedOn        time.Time `json:"updated_on,omitempty"`
}

func BQWriteTest() {
	transformDomainName("webmail.58ih8jhbt6ac.wpeproxy.com")
	transformDomainName("58ih8jhbt6ac.wpeproxy.com")
	transformDomainName("58ih8jhbt6ac.wpeproxy.com:8080")
	transformDomainName("wpeproxy.com")
}

func transformDomainName(d string) {
	var domainName string
	//var isPresent bool
	formatDomain := strings.Split(d, ":")
	match, _ := regexp.MatchString("wpeproxy", formatDomain[0])
	if match {
		secureSubDomain := strings.Split(formatDomain[0], ".")
		// Check if secureSubDomain is available in cache
		//domainName := "mytestdomainname"
		isPresent := false
		if !isPresent {
			// remove this logic once validated
			if strings.EqualFold(secureSubDomain[0], "wpeproxy") {
				domainName = "mytestdomainname"
			} else {
				domainName = secureSubDomain[0]
			}
		}
	} else {
		domainName = formatDomain[0]
	}
	fmt.Println("domainName: ", domainName)
}
