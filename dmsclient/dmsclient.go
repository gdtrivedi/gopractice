package dmsclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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

// Client implements interacting with DMS API
type Client struct {
	httpClient *http.Client // TODO: remove this it is not used.
	dmsAuthKey string
	url        string
	httpu      httputil.HTTPUtilities
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

	installJson, err := json.Marshal(install)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("installJson: ", string(installJson))
	}

	b := new(bytes.Buffer)
	if err := json.NewEncoder(b).Encode(install); err != nil {
		return false, err
	}
	//req, _ := http.NewRequest(http.MethodPut, dmsURL.String(), b)
	fmt.Println("dmsURL.String():", dmsURL.String())
	req, _ := http.NewRequest(http.MethodPut, dmsURL.String(), bytes.NewBuffer(installJson))
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
