package base64

import (
	"encoding/base64"
	"fmt"
)

func Base64Test() {
	str := "{\"url\":\"https://api2.wpenginedev.com/install/test2020dns6\",\"action\":\"update\",\"syntheticSource\":\"dns-records-service-prober\"}"
	encodedStr := encodeStr(str)
	fmt.Println("Encoded: ", encodedStr)
	decodedStr := decodeStr(encodedStr)
	fmt.Println("Decoded: ", decodedStr)
	input := []byte("eyJ1cmwiOiJodHRwczovL2FwaTIud3BlbmdpbmVkZXYuY29tL2luc3RhbGwvdGVzdDIwMjBkbnM2IiwiYWN0aW9uIjoidXBkYXRlIiwic3ludGhldGljU291cmNlIjoiZG5zLXJlY29yZHMtc2VydmljZS1wcm9iZXIifQ")
	fmt.Println("Decoded Bytes: ", decodeBytes(input))
}

func encodeStr(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func decodeStr(str string) string {
	decoded, _ := base64.StdEncoding.DecodeString(str)
	return string(decoded)
}
func decodeBytes(bytes []byte) string {
	//var dst []byte
	dst := make([]byte, base64.StdEncoding.DecodedLen(len(bytes)))
	base64.StdEncoding.Decode(dst, bytes)
	return string(dst)
}
