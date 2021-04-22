package jsontomap

import (
	"encoding/json"
	"fmt"
	"path/filepath"
)

func dumpMap(space string, m map[string]interface{}) {
	for k, v := range m {
		if mv, ok := v.(map[string]interface{}); ok {
			fmt.Printf("{ \"%v\": \n", k)
			dumpMap(space+"\t", mv)
			fmt.Printf("}\n")
		} else {
			fmt.Printf("%v %v : %v\n", space, k, v)
		}
	}
}

var jsonStr = `
{
  "array": [
	1,
	2,
	3
  ],
  "boolean": true,
  "null": null,
  "number": 123,
  "object": {
	"a": "b",
	"c": "d",
	"e": "f"
  },
  "string": "Hello World"
}
`

var jsonStr2 = `
{
    "success": {
        "uploadcert02.gtrivedi.xyz.key": "c3101f12cd2559bb0557f36ff2d115fc",
        "uploadcert02.gtrivedi.xyz.crt": "edbf51d92476f60c0d72fc1dc2939d3c"
    }
}
`

func JsontoMapTest() {
	jsonMap := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}
	dumpMap("", jsonMap)
	fmt.Println(jsonMap["string"])
	fmt.Println(jsonMap["object"].(map[string]interface{})["a"])
	arr := jsonMap["array"].([]interface{})
	fmt.Println(arr[0])

	jsonMap2 := make(map[string]interface{})
	err2 := json.Unmarshal([]byte(jsonStr2), &jsonMap2)
	if err2 != nil {
		panic(err)
	}
	fmt.Println(jsonMap2["success"])
	str := "/Users/gautam.trivedi/Downloads/upload_cert/uploadcert02.gtrivedi.xyz.crt"
	fmt.Println(jsonMap2["success"].(map[string]interface{})[filepath.Base(str)])
	fmt.Println(jsonMap2["success"].(map[string]interface{})[filepath.Base(str)].(string))
	fmt.Println(jsonMap2["success"].(map[string]interface{})["uploadcert02.gtrivedi.xyz.key"])
}
