package model

import (
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
)

type (
	// CreateCustomHostnameRequest represents a request to get a custom hostname
	CreateCustomHostnameRequest struct {
		CFZoneID           string `json:"cf_zone_id"`
		Hostname           string `json:"hostname"`
		CustomOriginServer string `json:"custom_origin_server"`
	}

	// ListCustomHostnameRequest represents a request to list custom hostnames
	ListCustomHostnameRequest struct {
		Hostname           *string
		CustomOriginServer *string
		AgentStatusState   *string
		SpecRemoved        *bool
		StatusRemoved      *bool
		Reconciled         *bool
		Limit              *uint64
		Offset             *uint64
	}

	// UpdateSpecRequest represents a request body for updating spec
	UpdateSpecRequest struct {
		CFZoneID           *string    `json:"cf_zone_id,omitempty"`
		Hostname           *string    `json:"hostname,omitempty"`
		CustomOriginServer *string    `json:"custom_origin_server,omitempty"`
		Removed            *bool      `json:"removed,omitempty"`
		UpdatedOn          *time.Time `json:"updated_on,omitempty"`
	}

	// UpdateStatusRequest represents a request body for updating status
	UpdateStatusRequest struct {
		CFZoneID                   *string           `json:"cf_zone_id,omitempty"`
		CFHostnameUUID             *string           `json:"cf_hostname_uuid,omitempty"`
		Hostname                   *string           `json:"hostname,omitempty"`
		CFHostnameStatus           *CFHostnameStatus `json:"cf_hostname_status,omitempty"`
		SSLStatus                  *string           `json:"ssl_status"`
		HostnameVerificationErrors []string          `json:"hostname_verification_errors"`
		CustomOriginServer         *string           `json:"custom_origin_server,omitempty"`
		Removed                    *bool             `json:"removed,omitempty"`
		UpdatedOn                  *time.Time        `json:"updated_on,omitempty"`
		SSLValidationRequestedOn   *time.Time        `json:"ssl_validation_requested_on,omitempty"`
	}

	// UpdateAgentStatusRequest represents a request body for updating agent status
	UpdateAgentStatusRequest struct {
		CustomHostnameUUID *uuid.UUID `json:"custom_hostname_uuid,omitempty"`
		State              *string    `json:"state,omitempty"`
		StateCount         *int32     `json:"state_count,omitempty"`
		Message            *string    `json:"message,omitempty"`
		UpdatedOn          *time.Time `json:"updated_on,omitempty"`
	}

	// UpdateRequest represents a request entity for updating customhostname record
	UpdateRequest struct {
		Spec        UpdateSpecRequest        `json:"spec"`
		Status      UpdateStatusRequest      `json:"status"`
		AgentStatus UpdateAgentStatusRequest `json:"agent_status"`
	}

	// CustomHostnameRequestEntity represents a request entity for fetching customhostname record
	CustomHostnameRequestEntity struct {
		UUID               uuid.UUID `json:"uuid"`
		Hostname           *string   `json:"hostname"`
		CustomOriginServer *string   `json:"custom_origin_server"`
	}

	// TriggerValidationRequest represents a request body to trigger Cloudflare validation
	TriggerValidationRequest struct {
		Hostname *string `json:"hostname,omitempty"`
		Install  *string `json:"install,omitempty"`
	}

	// CachePurge represents a request body to purge cache on Cloudflare
	CachePurge struct {
		Hostnames []string `json:"hostnames,omitempty"`
		Install   *string  `json:"install,omitempty"`
	}
)

// ToQuery generates url query parameters
func (r *ListCustomHostnameRequest) ToQuery() url.Values {
	m := url.Values{}

	if v := r.Hostname; v != nil {
		m.Add("spec_hostname", *v)
	}

	if v := r.CustomOriginServer; v != nil {
		m.Add("spec_custom_origin_server", *v)
	}

	if v := r.AgentStatusState; v != nil {
		m.Add("agent_status_state", *v)
	}

	if v := r.SpecRemoved; v != nil {
		m.Add("spec_removed", strconv.FormatBool(*v))
	}

	if v := r.StatusRemoved; v != nil {
		m.Add("status_removed", strconv.FormatBool(*v))
	}

	if v := r.Reconciled; v != nil {
		m.Add("reconciled", strconv.FormatBool(*v))
	}

	if v := r.Limit; v != nil {
		m.Add("limit", strconv.FormatUint(*v, 10))
	}

	if v := r.Offset; v != nil {
		m.Add("offset", strconv.FormatUint(*v, 10))
	}

	return m
}
