package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println(".....Start.....")

	file1Bytes := readFile("github.com/gdtrivedi/gopractice/jsonfiletostruct/settings1.json")
	var settings1 Settings
	json.Unmarshal(file1Bytes, &settings1)

	file2Bytes := readFile("github.com/gdtrivedi/gopractice/jsonfiletostruct/settings2.json")
	var settings2 Settings
	json.Unmarshal(file2Bytes, &settings2)

	fmt.Printf("String Comparision: %d", strings.Compare(settings1.String(), settings2.String()))
	fmt.Println()
	fmt.Printf("Equality (without ignoring any field): %t", cmp.Equal(settings1, settings2))
	fmt.Println()
	fmt.Printf("Equality (ignoring field): %t", cmp.Equal(settings1, settings2, cmpopts.IgnoreFields(Wifi{}, "Country")))
	fmt.Println()
	fmt.Println(".....End.....")
}

func readFile(filePath string) []byte {
	// Open our jsonFile
	jsonFile, err := os.Open(filePath)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Successfully Opened %s\n", filePath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	return byteValue
}