package split

import (
	"fmt"
	"strings"
)

func SplitTest() {
	fmt.Println("SSD: ", getSSD(strings.Split("abc.32sdi34sf.wpeproxy.com", ".")))
	fmt.Println("SSD: ", getSSD(strings.Split("abc.xyz.32sdi34sf.wpeproxy.com", ".")))
	fmt.Println("SSD: ", getSSD(strings.Split("32sdi34sf.wpeproxy.com", ".")))
	fmt.Println("SSD: ", getSSD(strings.Split("32sdi34sf.wpeproxy.co.uk", ".")))
	fmt.Println("SSD: ", getSSD(strings.Split("abc.32sdi34sf.wpeproxy.co.uk", ".")))
	urlParts(strings.Split("https://dms-staging.wpengine.io/dms/v1/domains/www.raol.eu", "/"))
}

func getSSD(secureSubDomain []string) string {
	for index, ssd := range secureSubDomain {
		if ssd == "wpeproxy" {
			return secureSubDomain[index-1]
		}
	}
	return ""
}

func urlParts(urlParts []string) {
	fmt.Println(urlParts[len(urlParts)-1])
}
