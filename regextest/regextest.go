package main

import (
	"fmt"
	"regexp"
)

func main() {
	serialNumber := "34110003T0122"
	//check if serial number is valid
	alphaNumericRegex := regexp.MustCompile("^[a-zA-Z0-9]*$")
	fmt.Println("Matched: ", alphaNumericRegex.MatchString(serialNumber))

}