package strconv

import (
	"fmt"
	"strconv"
	"strings"
)

// AllowedTLSVersions is the list of allowed TLS Version for Custom Hostname Records.
var AllowedTLSVersions []string = []string{"1.0", "1.1", "1.2", "1.3"}

func ParseFloatTest() {
	str := "1.3"
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println("Error converting string to float: ", err)
	} else {
		fmt.Println("Value: ", val)
	}

	if val != 1.1 || val != 1.2 {
		fmt.Println("Wrong Value: ", val)
	}

	fmt.Printf("invalid ssl_min_tls_version %s. Value must be from [%s].", str, strings.Join(AllowedTLSVersions[:], ", "))
}
