package firestore

import (
	"context"
	"encoding/json"
	"fmt"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

type Config struct {
	Properties []Prp `firestore:"properties"`
}

type Prp struct {
	Key   string `firestore:"key"`
	Value string `firestore:"value"`
}

func PrintDocumentJSON(collection string, document string) {
	ctx := context.Background()

	clientOption := option.WithCredentialsFile("/Users/gautam.trivedi/Documents/Personal/GCP_Keys/gcp-practice-288220.json")

	client, err := firestore.NewClient(ctx, "gcp-practice-288220", clientOption)
	if err != nil {
		fmt.Println("ERROR >>> ", err)
	}
	defer client.Close()

	var configRes Config
	fi_app, err := client.Collection(collection).Doc(document).Get(ctx)
	if err != nil {
		fmt.Println("ERROR >>> ", err)
	}
	fi_app.DataTo(&configRes)
	fmt.Println("------------- ARRAY -------------")
	printConfig(configRes)
	configMap := mapConfig(configRes)
	fmt.Println("------------- MAP -------------")
	for key, value := range configMap {
		fmt.Println("Key: ", key)
		fmt.Println("Value: ", value)
	}

	configMapJson, err := json.Marshal(configMap)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("------------- JSON -------------")
	fmt.Println("Config JSON: ", string(configMapJson))
}

func printConfig(config Config) {
	for _, prop := range config.Properties {
		fmt.Println("Key: ", prop.Key)
		fmt.Println("Value: ", prop.Value)
	}
}
func mapConfig(config Config) map[string]interface{} {
	props := make(map[string]interface{})
	for _, prop := range config.Properties {
		props[prop.Key] = prop.Value
	}
	return props
}
