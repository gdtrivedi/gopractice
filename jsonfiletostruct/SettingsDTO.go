package main

import (
	"fmt"
)

type Settings struct {
	Id string `json:"_id"`
	BearerPriorities []string `json:"bearer-priority"`
	Version Version `json:"version"`
	Wifi Wifi `json:"wifi"`
}

func (settings Settings) String() string {

	return fmt.Sprintf(
		"%s%s%s%s",
		settings.Id,
		settings.BearerPriorities,
		settings.Version,
		settings.Wifi)
}

type Version struct {
	ConfigSchema string `json:"configSchema"`
	LRUSoftware string `json:"lruSoftware"`
}

func (version Version) String() string {
	return fmt.Sprintf("%s%s",version.ConfigSchema, version.LRUSoftware)
}

type Wifi struct {
	Country string `json:"country"`
	GeoLocation bool `json:"geolocation"`
	Networks []Network `json:"networks"`
	Radios []Radio `json:"radios"`
}

func (wifi Wifi) String() string {
	return fmt.Sprintf("%s%t%s%s", wifi.Country, wifi.GeoLocation,
		wifi.Networks, wifi.Radios)
}

type Network struct {
	Id int `json:"_id"`
	Bearers []string `json:"bearers"`
	Hidden bool `json:"hidden"`
	Password string `json:"password"`
	Radio int `json:"radio"`
	SSID string `json:"ssid"`
	VVIP bool `json:"vvip"`
}

func (network Network) String() string {
	return fmt.Sprintf("%d%s%t%s%d%s%t",
		network.Id, network.Bearers, network.Hidden, network.Password,
		network.Radio, network.SSID, network.VVIP)
}

type Radio struct {
	Id int `json:"_id"`
	Band string `json:"band"`
	Cabin bool `json:"cabin"`
	Channel int `json:"channel"`
	Cloud bool `json:"cloud"`
	Name string `json:"name"`
	Radio int `json:"radio"`
	Txpower string `json:"txpower"`
}

func (radio Radio) String() string {
	return fmt.Sprintf("%d%s%t%d%t%s%d%s",
		radio.Id, radio.Band, radio.Cabin, radio.Channel, radio.Cloud,
		radio.Name, radio.Radio, radio.Txpower)
}