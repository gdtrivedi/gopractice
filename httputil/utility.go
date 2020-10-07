package httputil

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// HTTPUtilities interface that hold helper methods of the http util package
type HTTPUtilities interface {
	MakeRequest(r *http.Request) (*http.Response, error)
}

// Utility is struct for holding commonly used http utility helper methods
type Utility struct {
	httpClient *http.Client
}

// NewUtility is construct for Utility
func NewUtility(httpClient *http.Client) *Utility {
	return &Utility{
		httpClient: httpClient,
	}
}

// SetClientDefaults sets some sensible defaults on HTTP Clients
func SetClientDefaults(client *http.Client) *http.Client {
	if client == nil {
		client = &http.Client{}
	}
	if client.Timeout == 0 {
		client.Timeout = 180 * time.Second
	}
	return client
}

// MakeRequest helper method to make http calls to GET methods
func (u *Utility) MakeRequest(r *http.Request) (*http.Response, error) {

	var errorResponseBody string

	// making a http call
	res, err := u.httpClient.Do(r)
	if err != nil {
		return res, err
	}

	// parse error response body in case of non-2xx's
	if !(res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices) {
		b, _ := ioutil.ReadAll(res.Body)
		errorResponseBody = string(b)
	}

	switch {
	case res.StatusCode >= http.StatusOK && res.StatusCode < http.StatusMultipleChoices:
	case res.StatusCode >= http.StatusBadRequest && res.StatusCode < http.StatusInternalServerError:
		return res, errors.Errorf("HTTP status %d: client errors : error response body %s",
			res.StatusCode, errorResponseBody)
	case res.StatusCode >= http.StatusInternalServerError:
		return res, errors.Errorf("HTTP status %d: server errors : error response body %s",
			res.StatusCode, errorResponseBody)
	default:
		return res, errors.Errorf("HTTP status %d: error response body %s", res.StatusCode, errorResponseBody)
	}

	return res, nil

}
