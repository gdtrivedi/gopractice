package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	timeInRFC3339Format :=currentTime.Format(time.RFC3339)

	fmt.Println("timestamp: %s", timestamp)
	fmt.Println("timeInRFC3339Format: %s", timeInRFC3339Format)
}