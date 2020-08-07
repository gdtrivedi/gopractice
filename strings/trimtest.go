package strings

import (
	"fmt"
	"strings"
)

func TrimString() {
	name := "  Trim Test  "
	strings.TrimSpace(name)
	fmt.Println(name)
	fmt.Println(len(name))
	name = strings.TrimSpace(name)
	fmt.Println(name)
	fmt.Println(len(name))
}