package modeltest

import (
	"time"

	"github.com/google/uuid"
)

type (
	CustomHostname struct {
		UUID       uuid.UUID `json:"uuid"`
		Reconciled bool      `json:"reconciled"`
		Spec       Spec      `json:"spec"`
	}
	// Spec is the record's spec.
	Spec struct {
		CFZoneID           string      `json:"cf_zone_id"`
		Hostname           string      `json:"hostname"`
		CustomOriginServer string      `json:"custom_origin_server"`
		SSL                SSLSettings `json:"ssl"`
		Removed            bool        `json:"removed"`
		UpdatedOn          time.Time   `json:"updated_on"`
	}
	// SSLSettings is the record's SSL settings.
	SSLSettings struct {
		MinTLSVersion string `json:"min_tls_version,omitempty"`
	}
)
