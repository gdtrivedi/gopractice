package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
}

type LRULogRecord struct {
	Acid             string                 `json:"acid"`
	MsgCode          string                 `json:"msgCode"`
	HardwareRevision string                 `json:"hardwareRevision,omitempty"`
	SoftwareRevision string                 `json:"softwareRevision"`
	PartNumber       string                 `json:"partNumber"`
	Description      string                 `json:"description"`
	Time             int64                  `json:"time"`
	SerialNumber     string                 `json:"serialNumber"`
	SubSystem        string                 `json:"subSystem"`
	Parameters       map[string]interface{} `json:"parameters"`
	Metadata         map[string]interface{} `json:"metadata"`
}

func main() {
	user := &User{Name: "Frank"}
	b, err := json.Marshal(user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

	lru := LRULogRecord{
		Acid:             "AB1233",
		MsgCode:          "063",
		HardwareRevision: "3",
		SoftwareRevision: "4",
		PartNumber:       "P123",
		Description:      "description",
		Time:             1234567890,
		SerialNumber:     "4321",
		SubSystem:        "subSys",
		Parameters: map[string]interface{}{
			"id1": "val1",
			"id2": "val2",
		},
		Metadata: map[string]interface{}{
			"wow": "429",
			"anotherwow": map[string]interface{}{
				"wowo": "wowo",
			},
		},
	}

	strFormat := "Values of LRU Log Record:= ACID: %s, MsgCode:%s, HardwareRevision:%s, SoftwareRevision:%s, PartNumber:%s, " +
		"Description: %s, Time:%d, SerialNumber:%s, SubSystem:%s, Parameters:%v, Metadata:%v\n"
	fmt.Printf(strFormat, lru.Acid, lru.MsgCode, lru.HardwareRevision,
		lru.SoftwareRevision, lru.PartNumber, lru.Description,
		lru.Time, lru.SerialNumber, lru.SubSystem, lru.Parameters,
		lru.Metadata)

	params, ok := lru.Metadata["anotherwow"].(map[string]interface{})
	if ok {
		fmt.Println("wowo: ", params["wowo"])
	}

	lruJson, err := json.Marshal(lru)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("LRU Record JSON: ", string(lruJson))
}
