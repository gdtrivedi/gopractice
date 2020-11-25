package model

import "net/http"

// PageResponse is a response containing a set of CustomHostnameRecords.
type PageResponse struct {
	*Page
}

// Render renders a CustomHostnameRecord to the http.ResponseWriter.
func (r *CustomHostnameRecord) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// Render renders a CustomHostnameSpec to the http.ResponseWriter.
func (r *CustomHostnameSpec) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// Render renders a CustomHostnameStatus to the http.ResponseWriter.
func (r *CustomHostnameStatus) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// Render renders a AgentStatus to the http.ResponseWriter.
func (r *AgentStatus) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// Render renders a Response to the http.ResponseWriter.
func (r *CustomHostnameServiceInfo) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// Render renders a PageResponse to the http.ResponseWriter.
func (r *PageResponse) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// ValidationResponse is a struct that represents the server status object returned by the status endpoint
type ValidationResponse struct {
	Success     bool   `json:"success"`
	ValidatedOn string `json:"validated_on"`
}

// Render renders a Response to the http.ResponseWriter.
func (r *ValidationResponse) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}

// CachePurgeResponse is a struct that represents response object
type CachePurgeResponse struct {
	Success    bool   `json:"success"`
	ExecutedOn string `json:"executed_on"`
}

// Render renders a Response to the http.ResponseWriter.
func (r *CachePurgeResponse) Render(w http.ResponseWriter, req *http.Request) error {
	return nil
}
