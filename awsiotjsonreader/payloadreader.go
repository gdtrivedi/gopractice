package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	file1Bytes, err := readFile("github.com/gdtrivedi/gopractice/awsiotjsonreader/dhcp_subscription.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return
	}

	var payloadReader PayloadReader
	json.Unmarshal(file1Bytes, &payloadReader)

	//fmt.Println(payloadReader.Messages)

	for _, msg := range payloadReader.Messages {
		//fmt.Println(msg.Format)
		//fmt.Println(msg.Topic)
		// 3. Split by substr and check len of the slice, or length is 1 if substr is not present
		ss := strings.Split(msg.Topic, "/")
		fmt.Println("Tid: ", ss[2])
		fmt.Println("Timestamp: ", msg.Payload.Timestamp)

		dhcpJsonArr := []string{}
		for _, payloadEle := range msg.Payload.Payload {
			//fmt.Println("Key:", payloadKey, "=>", "Element:", payloadEle)

			//fmt.Println(reflect.TypeOf(payloadEle))

			for _, element := range payloadEle.([]interface{}) {

				dhcp := DHCPPayload{
					Tid:       ss[2],
				}

				timestampInt64, err := strconv.ParseInt(msg.Payload.Timestamp, 10, 64)
				if err != nil {
					continue
				}
				dhcp.EventTimestamp = timestampInt64

				dhcp.InsertTimestamp = time.Now().Unix()

				if val, ok := element.(map[string]interface{})["client"]; ok {
					dhcp.Client = val.(string)
				}

				if val, ok := element.(map[string]interface{})["connect"]; ok {
					dhcp.Connect = val.(string)
				}

				if val, ok := element.(map[string]interface{})["device"]; ok {
					dhcp.Device = val.(string)
				}

				if val, ok := element.(map[string]interface{})["expires"]; ok {
					dhcp.Expires = val.(string)
				}

				if val, ok := element.(map[string]interface{})["hostname"]; ok {
					dhcp.Hostname = val.(string)
				}

				if val, ok := element.(map[string]interface{})["ip"]; ok {
					dhcp.Ip = val.(string)
				}

				if val, ok := element.(map[string]interface{})["mac"]; ok {
					dhcp.Mac = val.(string)
				}

				if val, ok := element.(map[string]interface{})["ssid"]; ok {
					dhcp.Ssid = val.(string)
				}


				dhcpJson, err := json.Marshal(dhcp)
				if err != nil {
					fmt.Println(err)
					return
				}
				dhcpJsonArr = append(dhcpJsonArr, string(dhcpJson))
				//fmt.Println("DHCP: ", dhcp)
				//fmt.Println("IP: ", element.(map[string]interface{})["ip"],", Client: ", element.(map[string]interface{})["client"],", SSID: ", element.(map[string]interface{})["ssid"])
			}
			//var payloadEleArr = payloadEle.(map[string]interface{})
			//fmt.Println("payloadEleArr: ", payloadEleArr)
			//for key, ele := payloadEleArr {
			//	fmt.Println("Key:", key, "=>", "Element:", ele)
			//}
		}
		fmt.Println("DHCP Record JSON: ", dhcpJsonArr)

		if(len(dhcpJsonArr) > 0) {
			err := writeFile("github.com/gdtrivedi/gopractice/awsiotjsonreader/dhcp_" + ss[2] + "_" + msg.Payload.Timestamp +".json", dhcpJsonArr)
			// if we os.Open returns an error then handle it
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}

}

func readFile(filePath string) ([]byte, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Printf("Successfully Opened %s\n", filePath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue, nil
}

func writeFile(filePath string, linesToWrite []string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return err
	}

	datawriter := bufio.NewWriter(file)

	for _, line := range linesToWrite {
		_, _ = datawriter.WriteString(line + "\n")
	}

	datawriter.Flush()
	file.Close()

	return nil
}