package auth

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"

	"golang.org/x/oauth2"
)

const clientTimeout = 60 * time.Second

// BasicAuthConfig is a struct used to hold the basic-auth login for a BasicAuthHTTPClient
type BasicAuthConfig struct {
	URL      string
	Username string
	Password string
}

// tokenSorce is a token source for basic auth endpoints
type tokenSource struct {
	basicAuthConfig BasicAuthConfig
	httpClient      *http.Client
}

// NewBasicAuthHTTPClient creates either a BasicAuthHTTPClient or a default http.Client depending on whether or not
// basic auth config is provided.
func NewBasicAuthHTTPClient(basicAuthConfig BasicAuthConfig) (*http.Client, error) {

	if basicAuthConfig.URL == "" || basicAuthConfig.Username == "" || basicAuthConfig.Password == "" {
		return nil, errors.New("Missing basic-auth configuration for BasicAuthHTTPClient")
	}

	tokenSource := &tokenSource{
		basicAuthConfig: basicAuthConfig,
		httpClient:      &http.Client{Timeout: 180 * time.Second},
	}

	client := oauth2.NewClient(context.Background(), tokenSource)
	client.Timeout = clientTimeout
	return client, nil
}

// Token uses the basicAuthConfig to generate a token from whatever auth service endpoint is provided as URL in
// basicAuthConfig.
func (s *tokenSource) Token() (*oauth2.Token, error) {

	req, err := http.NewRequest(http.MethodPost, s.basicAuthConfig.URL, nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(s.basicAuthConfig.Username, s.basicAuthConfig.Password)

	res, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	var responseJSON map[string]interface{}
	err = json.Unmarshal(bodyBytes, &responseJSON)
	if err != nil {
		return nil, err
	}

	return s.retrieveTokenFromJSON(responseJSON)
}

func (s *tokenSource) retrieveTokenFromJSON(m map[string]interface{}) (*oauth2.Token, error) {
	token, ok := m["token"].(string)
	if !ok {
		return nil, errors.New("token not found in JSON response")
	}
	exp, ok := m["expires_on"].(string)
	if !ok {
		return nil, errors.New("expires_on not found in JSON response")
	}
	expTime, err := time.Parse(time.RFC3339, exp)
	expTimeUTC := expTime.UTC()
	if err != nil {
		return nil, err
	}
	return &oauth2.Token{
		AccessToken: token,
		Expiry:      expTimeUTC,
		TokenType:   "Token", // Required for Basic Auth
	}, nil
}
