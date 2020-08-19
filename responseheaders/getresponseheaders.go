package responseheaders

import (
	"fmt"
	"net/http"
	"net/url"
)

func GetResponseHeaders() {
	scheme := "http"
	domain := "cachetagtest.wptestinstall.com"
	wpeCdnCheck := "-wpe-cdncheck-"
	u := url.URL{
		Scheme: scheme,
		Host:   domain,
		Path:   wpeCdnCheck,
	}

	// Prep http request
	req, err := http.NewRequest("HEAD", u.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "WPE-DSS-1.0")

	// Execute http request
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	//wpeEdgeHeader := "X-WPE-Edge"
	wpeEdgeHeader := "x-wpe-edge" //case of letters do not matter for headers.

	headers := resp.Header
	for name, value := range headers {
		fmt.Printf("%v: %v\n", name, value)
	}
	fmt.Println("Header >>> ", headers.Get(wpeEdgeHeader))
}
