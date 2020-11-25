package modeltest

import (
	"fmt"
	"time"

	"github.com/cloudflare/cloudflare-go"

	"github.com/google/uuid"
)

func ModelTest() {
	chr1 := &CustomHostname{
		UUID:       uuid.UUID{},
		Reconciled: false,
	}

	chr2 := &CustomHostname{
		UUID:       uuid.UUID{},
		Reconciled: false,
		Spec: Spec{
			CFZoneID:           "",
			Hostname:           "",
			CustomOriginServer: "",
			SSL:                SSLSettings{MinTLSVersion: "1.2"},
			Removed:            false,
			UpdatedOn:          time.Time{},
		},
	}

	fmt.Println("SSL:", chr1.Spec.SSL.MinTLSVersion)
	fmt.Println("SSL:", chr2.Spec.SSL.MinTLSVersion)

	if chr1.Spec.SSL.MinTLSVersion != chr2.Spec.SSL.MinTLSVersion {
		fmt.Println("Not Equal")
	} else {
		fmt.Println("Equal")
	}

	ch := cloudflare.CustomHostname{
		ID:                 "",
		Hostname:           "",
		CustomOriginServer: "",
		//SSL:                       cloudflare.CustomHostnameSSL{},
		CustomMetadata:            nil,
		Status:                    "",
		VerificationErrors:        nil,
		OwnershipVerification:     cloudflare.CustomHostnameOwnershipVerification{},
		OwnershipVerificationHTTP: cloudflare.CustomHostnameOwnershipVerificationHTTP{},
	}
	if ch.SSL.Settings.MinTLSVersion != chr2.Spec.SSL.MinTLSVersion {
		ch.SSL.Settings.MinTLSVersion = chr2.Spec.SSL.MinTLSVersion
		fmt.Println("CF SSL:", ch.SSL.Settings.MinTLSVersion)
	}
}
