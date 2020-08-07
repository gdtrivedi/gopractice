package jsonparsing

import (
	"encoding/json"
	"fmt"
)

type CachePurge struct {
	Hostnames []string `json:"hostnames,omitempty"`
	Install   *string  `json:"install,omitempty"`
}

type Message struct {
	Data []byte
}

func UnmarshallTest() {
	cp := &CachePurge{}
	m := `{"hostnames": [], "install": "abc"}`
	message := &Message{
		Data: []byte(m),
	}
	err := json.Unmarshal(message.Data, cp)
	if err != nil {
		fmt.Println("<<<<< ERROR >>>>>", err)
	} else {
		fmt.Println("Hostnames:", len(cp.Hostnames))
		fmt.Println("Install: ", *cp.Install)
	}
}
