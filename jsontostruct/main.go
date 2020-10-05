package jsontostruct

import (
	"encoding/json"
	"fmt"
)

type Child struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Parents []Parent `json:"parents"`
}

type Parent struct {
	Name string `json:"name"`
}

func main() {
	datamap := make(map[string]interface{})

	/*
				map[acid:1DDDDD description:Critical Test hardwareRevision:A id:1EEEFE_25_A_1.7_P2999 metadata:map[crmResponse:map[crmCaseId:5008A000003BWCqQAO success:true]] msgCode:063 parameters:map[wow:a429] partNumber:P2800 serialNumber:00030 softwareRevision:1.7 subSystem:SW]

				msg := &Message{
		        Action: "get_products",
		        Params: map[string]interface{}{
		            "id1": val1,
		            "id2": val2,
		        },
		    }
	*/
	datamap["acid"] = "1DDDDD"
	datamap["description"] = "Critical Test"
	datamap["hardwareRevision"] = "A"
	datamap["id"] = "1EEEFE_25_A_1.7_P2999"
	datamap["msgCode"] = "063"
	datamap["partNumber"] = "P2800"
	datamap["serialNumber"] = "00030"
	datamap["softwareRevision"] = "1.7"
	datamap["subSystem"] = "SW"

	// convert map to json
	jsonString, _ := json.Marshal(datamap)

	var child Child
	// json.NewDecoder().Decode(datamap)
	err := json.Unmarshal(jsonString, &child)

	if err != nil {
		fmt.Print("Error: ", err)
	}

	fmt.Println(child.Name)
}
