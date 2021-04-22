package dmsclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/gdtrivedi/gopractice/httputil"
	"github.com/pkg/errors"
)

// Install represents an install in the DMS service
type Install struct {
	URL     string          `json:"url"`
	Name    string          `json:"name"`
	ID      int             `json:"wpe_id"`
	Active  bool            `json:"active"`
	Domains []InstallDomain `json:"domains"`
}

// InstallDomain represents the domain model received from install endpoints
type InstallDomain struct {
	Name      string   `json:"name"`
	Dev       bool     `json:"dev"`
	Redirects []string `json:"redirects"`
}

// Domain represents a domain in the DMS service
type Domain struct {
	ID          *int32    `json:"id"`
	InstallID   *int32    `json:"install_id"`
	AtlasEnvID  string    `json:"atlas_environment_id"`
	Name        string    `json:"name"`
	Dev         *bool     `json:"dev"`
	RedirectTo  *int32    `json:"redirects_to"`
	InstallName string    `json:"install_name"`
	CreatedAt   time.Time `json:"created_at,string"`
	UpdatedAt   time.Time `json:"updated_at,string"`
}

// DomainResponse represents the full repose from DMS with all the URL and a list of domains
type DomainResponse struct {
	URL     string   `json:"url"`
	Domains []Domain `json:"domains"`
}

// AtlasEnvironment Holds the information for DMS Atlas Environment record.
type AtlasEnvironment struct {
	URL          string           `json:"url"`
	ID           string           `json:"env_id"`
	OriginServer string           `json:"origin_server"`
	Active       bool             `json:"active"`
	Domains      []AtlasEnvDomain `json:"domains"`
}

// AtlasEnvDomain represents the domain model received from atlas environment endpoints
type AtlasEnvDomain Domain

// Client implements interacting with DMS API
type Client struct {
	httpClient *http.Client // TODO: remove this it is not used.
	dmsAuthKey string
	url        string
	httpu      httputil.HTTPUtilities
}

// DomainListResponse represents the full repose from DMS with all the URL and a list of domains
type DomainListResponse struct {
	NextURL     string   `json:"next_url"`
	PreviousURL string   `json:"previous_url"`
	Domains     []Domain `json:"results"`
	Count       *int32   `json:"count"`
}

// NewClient constructs a RealClient
func NewClient(httpClient *http.Client, dmsAuthKey, baseURL string, httpu httputil.HTTPUtilities) *Client {

	return &Client{
		httpClient: httpClient,
		dmsAuthKey: dmsAuthKey,
		url:        baseURL,
		httpu:      httpu,
	}
}

// FetchInstallInfoByInstallName returns information about an install including the install name and list of domains
func (c *Client) FetchInstallInfoByInstallName(installName string) (*Install, error) {

	dmsInstall := &Install{}

	requestURI := fmt.Sprintf("%s/installs/%s", c.url, installName)
	dmsURL, err := url.ParseRequestURI(requestURI)
	if err != nil {
		return dmsInstall, err
	}

	req, _ := http.NewRequest(http.MethodGet, dmsURL.String(), nil)

	res, err := c.httpu.MakeRequest(c.addRequestHeaders(req))
	if err != nil {
		return dmsInstall, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&dmsInstall); err != nil {
		return dmsInstall, err
	}

	return dmsInstall, nil
}

// UpdateInstallInfo Updates the install with domains and corresponding redirects.
func (c *Client) UpdateInstallInfo(install *Install) (bool, error) {
	requestURI := fmt.Sprintf("%s/installs/%s", c.url, install.Name)
	dmsURL, err := url.ParseRequestURI(requestURI)
	if err != nil {
		return false, err
	}

	//installJson, err := json.Marshal(install)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("installJson: ", string(installJson))
	//}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(install); err != nil {
		return false, err
	}
	//req, _ := http.NewRequest(http.MethodPut, dmsURL.String(), b)
	//fmt.Println("dmsURL.String():", dmsURL.String())
	req, _ := http.NewRequest(http.MethodPut, dmsURL.String(), b)
	req.Header.Set("Content-Type", "application/json")
	res, err := c.httpu.MakeRequest(c.addRequestHeaders(req))
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	// Getting 204 No Content is successful call.
	if res.StatusCode != http.StatusNoContent {
		return false, errors.Errorf("Some error occurred while updating install: %s", install.Name)
	}

	return true, nil
}

// addRequestHeaders is helper method to append custom headers to the Request Object
func (c *Client) addRequestHeaders(r *http.Request) *http.Request {
	r.Header.Set("Authorization", "token "+c.dmsAuthKey)

	// Setting this field prevents re-use of TCP connections between requests to the same hosts, as if
	// Transport.DisableKeepAlives were set.
	r.Close = true

	return r
}

// GetDomainInfoByID Returns the domain information for the domain ID.
func (c *Client) GetDomainInfoByID(domainID *int32) (*Domain, error) {
	uriString := fmt.Sprintf("%s/domains/%d", c.url, *domainID)
	domains, err := c.retrieveDomains(uriString)
	if err != nil {
		return nil, err
	}

	if len(domains) <= 0 {
		return nil, nil
	}

	// As we are looking for domain by ID, there should always be 1 exactly.
	domain := &Domain{
		ID:          domains[0].ID,
		InstallID:   domains[0].InstallID,
		AtlasEnvID:  domains[0].AtlasEnvID,
		Name:        domains[0].Name,
		Dev:         domains[0].Dev,
		RedirectTo:  domains[0].RedirectTo,
		InstallName: domains[0].InstallName,
		CreatedAt:   domains[0].CreatedAt,
		UpdatedAt:   domains[0].UpdatedAt,
	}

	return domain, nil
}

func (c *Client) GetDomainsByName(domainName string) ([]Domain, error) {
	uriString := fmt.Sprintf("%s/domains?name=%s&active_backend=%v", c.url, domainName, true)
	domains, err := c.retrieveDomainsList(uriString)
	if err != nil {
		return nil, err
	}
	return domains, nil
}

func (c *Client) retrieveDomainsList(domainsURL string) ([]Domain, error) {
	var domains []Domain
	//body, err := c.retrieveDomainData(domainsURL)
	//
	//if body == nil {
	//	return domains, err
	//}
	//
	//dr, err := decodeDomainListResponse(body)
	//if err != nil {
	//	return domains, err
	//}

	dr, err := c.retrieveDomainDataWithResponse(domainsURL)
	if err != nil {
		return domains, err
	}
	return dr.Domains, nil
}

func (c *Client) retrieveDomainData(domainsURL string) (io.ReadCloser, error) {
	requestURL, err := url.ParseRequestURI(domainsURL)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodGet, requestURL.String(), nil)

	res, err := c.httpu.MakeRequest(c.addRequestHeaders(req))
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, err
	}

	defer res.Body.Close()
	return res.Body, nil
}

func (c *Client) retrieveDomainDataWithResponse(domainsURL string) (*DomainListResponse, error) {
	requestURL, err := url.ParseRequestURI(domainsURL)
	if err != nil {
		return nil, err
	}

	req, _ := http.NewRequest(http.MethodGet, requestURL.String(), nil)

	res, err := c.httpu.MakeRequest(c.addRequestHeaders(req))
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return nil, err
	}

	defer res.Body.Close()

	var dlr DomainListResponse
	if err := json.NewDecoder(res.Body).Decode(&dlr); err != nil {
		return nil, nil
	}

	return &dlr, err
}

func decodeDomainListResponse(body io.Reader) (DomainListResponse, error) {
	var dlr DomainListResponse
	err := json.NewDecoder(body).Decode(&dlr)
	return dlr, err
}

func (c *Client) retrieveDomains(domainsURL string) ([]Domain, error) {
	var domains []Domain
	requestURL, err := url.ParseRequestURI(domainsURL)
	if err != nil {
		return domains, err
	}

	req, _ := http.NewRequest(http.MethodGet, requestURL.String(), nil)

	res, err := c.httpu.MakeRequest(c.addRequestHeaders(req))
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return domains, nil
		}
		return domains, err
	}

	defer res.Body.Close()

	dr, err := decodeDomainResponse(res.Body)
	if err != nil {
		return domains, err
	}

	return dr.Domains, nil
}
func decodeDomainResponse(body io.Reader) (DomainResponse, error) {
	var dr DomainResponse
	err := json.NewDecoder(body).Decode(&dr)
	return dr, err
}

// FetchAtlasEnvInfoByID fetches atlas environment details by ID.
func (c *Client) FetchAtlasEnvInfoByID(atlasEnvID string) (*AtlasEnvironment, error) {
	atlasEnv := &AtlasEnvironment{}

	requestURI := fmt.Sprintf("%s/atlas_environments/%s", c.url, atlasEnvID)

	dmsURL, err := url.ParseRequestURI(requestURI)
	if err != nil {
		return atlasEnv, err
	}

	fmt.Println("dmsURL.String(): ", dmsURL.String())

	req, _ := http.NewRequest(http.MethodGet, dmsURL.String(), nil)

	res, err := c.httpu.MakeRequest(c.addRequestHeaders(req))
	if err != nil {
		if res != nil && res.StatusCode == http.StatusNotFound {
			return nil, nil
		}
		return atlasEnv, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&atlasEnv); err != nil {
		return atlasEnv, err
	}

	return atlasEnv, nil
}
