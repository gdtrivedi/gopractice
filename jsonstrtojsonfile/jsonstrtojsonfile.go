package main

import (
	"encoding/json"
	"fmt"
)

type DHCPPayload struct {
	Tid string `json:"tid"`
	//Client    string   `json:"client"`
	//Connect   string	`json:"connect"`
	Device string `json:"device"`
	//Expires string `json:"expires"`
	Hostname string `json:"hostname"`
	Ip string `json:"ip"`
	Mac string `json:"mac"`
	Ssid string `json:"ssid"`
}

var jsonStr = `
{
  "payload": {
    "Andrews-iPhone.gogo.aero": [
      {
        "client": "int-cabin",
        "connect": "false",
        "device": "wlan0",
        "expires": "2020/05/08 20:38:59",
        "hostname": "Andrews-iPhone.gogo.aero",
        "ip": "172.20.10.101",
        "mac": "90:e1:7b:b1:87:be",
        "ssid": "N87BCwifi5.0GHz"
      }
    ],
    "Capt-Dave.gogo.aero": [
      {
        "client": "int-cabin",
        "connect": "true",
        "device": "wlan1",
        "expires": "2020/05/08 20:39:12",
        "hostname": "Capt-Dave.gogo.aero",
        "ip": "172.20.10.107",
        "mac": "f4:06:16:52:b6:3e",
        "ssid": "N87BCwifi2.4GHz"
      }
    ],
    "DavidHenkesiPad.gogo.aero": [
      {
        "client": "int-cabin",
        "connect": "true",
        "device": "wlan0",
        "expires": "2020/05/08 20:38:31",
        "hostname": "DavidHenkesiPad.gogo.aero",
        "ip": "172.20.10.103",
        "mac": "5c:ad:cf:ac:1e:f7",
        "ssid": "N87BCwifi5.0GHz"
      }
    ],
    "Sarahs-iPhone.gogo.aero": [
      {
        "client": "int-cabin",
        "connect": "true",
        "device": "wlan0",
        "expires": "2020/05/08 20:38:52",
        "hostname": "Sarahs-iPhone.gogo.aero",
        "ip": "172.20.10.104",
        "mac": "44:4a:db:0f:73:b1",
        "ssid": "N87BCwifi5.0GHz"
      }
    ],
    "iPad-145.gogo.aero": [
      {
        "client": "int-cabin",
        "connect": "true",
        "device": "wlan0",
        "expires": "2020/05/08 20:38:43",
        "hostname": "iPad-145.gogo.aero",
        "ip": "172.20.10.102",
        "mac": "5c:97:f3:dd:92:66",
        "ssid": "N87BCwifi5.0GHz"
      }
    ]
  },
  "timestamp": "1588970301"
}`

func main() {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	fmt.Println("Timestamp: ", jsonMap["timestamp"])
	for _, payloadEle := range jsonMap["payload"].(map[string]interface{}) {
		//fmt.Println("Key:", payloadKey, "=>", "Element:", payloadEle)

		//fmt.Println(reflect.TypeOf(payloadEle))
		for _, element := range payloadEle.([]interface{}) {
			dhcp := DHCPPayload{
				Tid:      "",
				Device:   "",
				Hostname: "",
				Ip:       "",
				Mac:      "",
				Ssid:     "",
			}

			fmt.Println("Element: ", element)
			fmt.Println("Client: ", element.(map[string]interface{})["client"])
			fmt.Println("SSID: ", element.(map[string]interface{})["ssid"])
		}
		//var payloadEleArr = payloadEle.(map[string]interface{})
		//fmt.Println("payloadEleArr: ", payloadEleArr)
		//for key, ele := payloadEleArr {
		//	fmt.Println("Key:", key, "=>", "Element:", ele)
		//}
	}
}
