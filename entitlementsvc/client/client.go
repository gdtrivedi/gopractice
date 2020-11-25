package client

import (
	"fmt"

	"github.com/gdtrivedi/gopractice/entitlementsvc/pkg/auth"
)

func ProvideEntitlementClient(url, username, password string) {
	basicAuthConfig := auth.BasicAuthConfig{
		URL:      url,
		Username: username,
		Password: password,
	}
	httpClient, err := auth.NewBasicAuthHTTPClient(basicAuthConfig)
	if err != nil {
		fmt.Println("Error: ", err)
	} else if httpClient == nil {
		fmt.Println("httpClient Nil")
	} else {
		fmt.Println("httpClient NOT Nil")
	}
	fmt.Println("<<< HERE >>>")
}
