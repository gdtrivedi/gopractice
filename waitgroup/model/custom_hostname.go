package model

import (
	"time"

	"github.com/google/uuid"
)

// CFHostnameStatus represents state of Hostname-Verification-Status in Cloudflare
type CFHostnameStatus string

const (
	// HostnameStatusPending status represents state of Hostname-Verification-Status in Cloudflare is pending.
	HostnameStatusPending CFHostnameStatus = "pending"
	// HostnameStatusActive status represents state of Hostname-Verification-Status in Cloudflare is active.
	HostnameStatusActive CFHostnameStatus = "active"
	// HostnameStatusMoved status represents state of Hostname-Verification-Status in Cloudflare is moved.
	HostnameStatusMoved CFHostnameStatus = "moved"
	// HostnameStatusRemoved status represents state of Hostname-Verification-Status in Cloudflare is removed.
	HostnameStatusRemoved CFHostnameStatus = "removed"
	// HostnameStatusDeleted status represents state of Hostname-Verification-Status in Cloudflare is deleted.
	HostnameStatusDeleted CFHostnameStatus = "deleted"
	// HostnameStatusUnknown status represents default state of Hostname-Verification-Status in farm-ingress-api
	// and value to be set when CustomHostnameRecord.Status.removed = true.
	HostnameStatusUnknown CFHostnameStatus = "unknown"
)

type (
	// Page holds a list of CustomHostnameRecord structs, along with a count of the total number in the collection.
	Page struct {
		Count                 uint64                 `json:"count"`
		CustomHostnameRecords []CustomHostnameRecord `json:"results"`
		Next                  string                 `json:"next"`
	}

	// CustomHostnameRecord represents a single reconcilable record
	CustomHostnameRecord struct {
		ID          uint64               `json:"-"`
		UUID        uuid.UUID            `json:"uuid"`
		URL         string               `json:"url"`
		Spec        CustomHostnameSpec   `json:"spec"`
		Status      CustomHostnameStatus `json:"status"`
		AgentStatus AgentStatus          `json:"agent_status"`
		Reconciled  bool                 `json:"reconciled"`
		CreatedOn   time.Time            `json:"created_on"`
	}

	// CustomHostnameSpec is the record's spec
	CustomHostnameSpec struct {
		CFZoneID           string    `json:"cf_zone_id"`
		URL                string    `json:"url"`
		Hostname           string    `json:"hostname"`
		CustomOriginServer string    `json:"custom_origin_server"`
		Removed            bool      `json:"removed"`
		UpdatedOn          time.Time `json:"updated_on"`
	}

	// CustomHostnameStatus is the record's status
	CustomHostnameStatus struct {
		URL                        string           `json:"url"`
		CFZoneID                   string           `json:"cf_zone_id"`
		CFHostnameUUID             string           `json:"cf_hostname_uuid"`
		Hostname                   string           `json:"hostname"`
		CFHostnameStatus           CFHostnameStatus `json:"cf_hostname_status"`
		SSLStatus                  string           `json:"ssl_status"`
		HostnameVerificationErrors []string         `json:"hostname_verification_errors"`
		CustomOriginServer         string           `json:"custom_origin_server"`
		Removed                    bool             `json:"removed"`
		UpdatedOn                  time.Time        `json:"updated_on"`
		SSLValidationRequestedOn   time.Time        `json:"ssl_validation_requested_on"`
	}

	// AgentStatus represents the reconciliation's agent status
	AgentStatus struct {
		UUID               uuid.UUID `json:"uuid"`
		URL                string    `json:"url"`
		CustomHostnameUUID uuid.UUID `json:"custom_hostname_uuid"`
		State              string    `json:"state"`
		StateCount         int       `json:"state_count"`
		Message            string    `json:"message"`
		CreatedOn          time.Time `json:"created_on"`
		UpdatedOn          time.Time `json:"updated_on"`
	}

	// CustomHostnameServiceInfo is a struct that represents the service details returned by the service info endpoint
	CustomHostnameServiceInfo struct {
		IPAddresses []string `json:"ip_addresses"`
		CNAME       string   `json:"cname"`
		RootDomain  string   `json:"root_domain"`
		CFZoneID    string   `json:"cf_zone_id"`
	}
)
