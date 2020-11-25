package waitgroup

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"sync"

	"github.com/gdtrivedi/gopractice/waitgroup/model"
	"google.golang.org/protobuf/proto"
)

const (
	baseURL                    = "https://farm-ingress-api-staging.wpesvc.net"
	path                       = "/v1/custom_hostnames"
	customHostNameOriginServer = "farmingressloa.wpenginepoweredstaging.com"
)

// Page holds a short list of entity structs, along with a count of the total number in the collection
type Page struct {
	Count    int                          `json:"count"`
	Entities []model.CustomHostnameRecord `json:"results"`
	Next     string                       `json:"Next"`
	// Skipping previous pagination fields for now
}

// HTTPEntityClient is a concrete entityClient that calls out to an API3 API
type HTTPEntityClient struct {
	Client *http.Client
}

func ReadCustomHostnamesTest() {
	token := "token of farm-ingress api"

	bearer_token, _ := fmt.Printf("Bearer %s", token) //"Bearer {}".format(token)

	fmt.Println("bearer_token: ", bearer_token)

	count := 50

	firstPage := true
	fetchAndProcessPages(count, firstPage)

}

func fetchAndProcessPages(count int, firstPage bool) {
	entityClient := &HTTPEntityClient{Client: &http.Client{}}
	var url string
	var waitgroup sync.WaitGroup
	for count > 0 || firstPage {
		if firstPage {
			url = getFirstPageURL()
			firstPage = false
		}
		page, _ := entityClient.FetchPage(url)
		url = page.Next
		count = page.Count
		waitgroup.Add(1)
		go processPage(page.Entities, &waitgroup)
	}
}
func processPage(chrs []model.CustomHostnameRecord, waitgroup *sync.WaitGroup) {
	for _, chr := range chrs {
		fmt.Println("CHR: ", chr.UUID)
	}
	waitgroup.Done()
}
func getFirstPageURL() string {
	r := &model.ListCustomHostnameRequest{
		CustomOriginServer: proto.String(customHostNameOriginServer),
		SpecRemoved:        proto.Bool(false),
		Reconciled:         proto.Bool(true),
		Limit:              proto.Uint64(20),
	}
	v := buildRemovedQuery(r)

	u, err := url.ParseRequestURI(baseURL + path)

	if err != nil {
		fmt.Println(err)
	}

	u.RawQuery = v.Encode()
	return u.String()
}
func buildRemovedQuery(r *model.ListCustomHostnameRequest) url.Values {
	v := url.Values{}
	v.Set("limit", strconv.FormatUint(*r.Limit, 10))

	if r.Offset != nil {
		v.Set("offset", strconv.FormatUint(*r.Offset, 10))
	}

	if r.Reconciled != nil {
		v.Set("reconciled", strconv.FormatBool(*r.Reconciled))
	}

	if r.Hostname != nil {
		v.Set("spec_hostname", *r.Hostname)
	}

	if r.SpecRemoved != nil {
		v.Set("spec_removed", strconv.FormatBool(*r.SpecRemoved))
	}

	if r.CustomOriginServer != nil {
		v.Set("spec_custom_origin_server", *r.CustomOriginServer)
	}

	return v
}

// FetchPage retrieves a page of entities
func (c *HTTPEntityClient) FetchPage(urlString string) (Page, error) {
	page := Page{}

	req, err := http.NewRequest("GET", urlString, nil)
	if err != nil {
		return page, err
	}

	res, err := c.MakeRequest(req)
	if err != nil {
		return page, err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&page)

	return page, err
}

// MakeRequest makes a HTTP request and returns the response
func (c *HTTPEntityClient) MakeRequest(r *http.Request) (*http.Response, error) {
	r.Close = true
	r.Header["Content-type"] = []string{"application/json"}
	r.Header["Accept"] = []string{"application/json"}

	//nolint:gosec
	token := "eyJhbGciOiJSUzI1NiIsImtpZCI6ImYwOTJiNjEyZTliNjQ0N2RlYjEwNjg1YmI4ZmZhOGFlNjJmNmFhOTEiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOiJodHRwczovL2Zhcm0taW5ncmVzcy1hcGktc3RhZ2luZy5lbmRwb2ludHMud3AtZW5naW5lLWNvcnBvcmF0ZS5jbG91ZC5nb29nIiwiYXpwIjoiZmFybS1pbmdyZXNzLWFwaS1zdGFnZUB3cC1lbmdpbmUtY29ycG9yYXRlLmlhbS5nc2VydmljZWFjY291bnQuY29tIiwiZW1haWwiOiJmYXJtLWluZ3Jlc3MtYXBpLXN0YWdlQHdwLWVuZ2luZS1jb3Jwb3JhdGUuaWFtLmdzZXJ2aWNlYWNjb3VudC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiZXhwIjoxNjA0NTM4NjEyLCJpYXQiOjE2MDQ1MzUwMTIsImlzcyI6Imh0dHBzOi8vYWNjb3VudHMuZ29vZ2xlLmNvbSIsInN1YiI6IjExNTM1ODI5MzYwMDc2NjgyNjM2MiJ9.gzEbds6UgI4TfCAxvDApKNV_ZE0qT1G8C_JjLkSXq17I4jbC8GC8tLE11svyfs7Oxt5x9QesIrdsSMXI4dnPXB5LZfKouAz7PIY2hugjmORa9UaJK7sfNi3x7qVcPJCChCKWbal-3y-DrlelMCNl_RP135-CKyvplHV5GrYMM3iTx3slNJhFtTAZRZuOyX_e0mbcjMvZ-jB8SY5dTTqzve4ANeEWrmJvqadNiCDgcqFDsdrjOnvVOrvfpQkwNRtcdJmjta-KabiQwgz8z8kRjZbyhu4C5kFbkysEnS6sFrivjlIZK-WRXtd7sKPiFhpI2UZfm2uZ7het7GZCi1iwuA"

	r.Header.Set("Authorization", "Bearer "+token)

	return c.Client.Do(r)
}
