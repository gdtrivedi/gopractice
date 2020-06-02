package main

import (
	"fmt"
	"strings"
)

func main() {
	name := "  Trim Test  "
	strings.TrimSpace(name)
	fmt.Println(name)
	fmt.Println(len(name))
	name = strings.TrimSpace(name)
	fmt.Println(name)
	fmt.Println(len(name))
}