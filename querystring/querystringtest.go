package querystring

import (
	"fmt"
	"net/url"

	"github.com/golang/protobuf/proto"
	"github.com/google/go-querystring/query"
	"github.com/pkg/errors"
)

const (
	EntitlementsResourcePath = "/v2/entitlements"
	// GESProductName is the name of the GI product in the entitlements service.
	GIProductName = "global-ingress-beta"
	// PlatformVendorName is the vendor name associated with platform api 1/2 and cmdb.
	PlatformVendorName = "wphosting"
)

// EntitlementsQuery represents the available parameters to query all Entitlements.
type EntitlementsQuery struct {
	AccountID      *int32  `url:"account_id,omitempty"`
	Vendor         *string `url:"vendor,omitempty"`
	SubscriptionID *int32  `url:"subscription_id,omitempty"`
	Product        *string `url:"product,omitempty"`
}

func QueryStringTest() {
	q := &EntitlementsQuery{
		AccountID:      proto.Int32(int32(1)),
		Product:        proto.String(GIProductName),
		Vendor:         proto.String(PlatformVendorName),
		SubscriptionID: proto.Int32(int32(2)),
	}

	u, err := entitlementsURL(q)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(u.String())
		fmt.Println(u.RequestURI())
	}
}

// EntitlementsURL generates url from EntitlementsQuery.
func entitlementsURL(q *EntitlementsQuery) (url.URL, error) {
	var u url.URL

	v, err := query.Values(q)
	if err != nil {
		return u, errors.Wrap(err, "invalid params")
	}

	baseURL, err := url.Parse("https://entitlements-staging.wpesvc.net")

	if err != nil {
		return u, err
	}

	u = *baseURL
	u.Path = EntitlementsResourcePath
	u.RawQuery = v.Encode()

	return u, nil
}
