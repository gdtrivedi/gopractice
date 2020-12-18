package timestamp

import (
	"fmt"
	"time"
)

func ShowTime() {
	currentTime := time.Now()
	timestamp := currentTime.Unix()
	timeInRFC3339Format := currentTime.Format(time.RFC3339)
	timeInStampMilliFormat := currentTime.Format(time.StampMilli)

	fmt.Println("timestamp: ", timestamp)
	fmt.Println("timeInRFC3339Format: ", timeInRFC3339Format)
	fmt.Println("timeInStampMilliFormat: ", timeInStampMilliFormat)

	time.Sleep(2 * time.Second)

	latency := time.Since(currentTime)
	fmt.Println("latency: ", latency)
}
