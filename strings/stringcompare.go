package strings

import (
	"fmt"
	"strings"
)

func StringCompare() {
	fmt.Println(strings.Compare("a", "b"))
	fmt.Println(strings.Compare("a", "A"))
	fmt.Println(strings.Compare("b", "a"))

	a := "gautam"
	b := "gautam"

	if a == b {
		fmt.Println("Equal")
	}

	// trim
	fmt.Println(len(strings.Trim("    Gautam   ", " ")))

}
