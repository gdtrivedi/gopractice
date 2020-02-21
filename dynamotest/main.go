package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

//SSA struct
type SSA struct {
	AccountGUID string    `json:"accountGuid"`
	SsaGUID     string    `json:"ssaGuid"`
	Aircraft    *Aircraft `json:"aircraft"`
}

//Aircraft struct
type Aircraft struct {
	AircraftGUID string `json:"aircraftGuid,omitempty"`
	Make         string `json:"make,omitempty"`
	Model        string `json:"model,omitempty"`
	Serial       string `json:"serial,omitempty"`
	PrimaryUse   string `json:"primaryUse,omitempty"`
	NewTail      string `json:"newTail,omitempty"`
	Tail         string `json:"tail,omitempty"`
}

func main() {
	fmt.Println("starting")

	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000")})
	if err != nil {
		log.Println(err)
		return
	}
	dbSvc := dynamodb.New(sess)

	result, err := dbSvc.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Tables:")
	for _, table := range result.TableNames {
		log.Println(*table)
	}

	fmt.Println("done")

	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String("DASHSSA"),
		KeyConditions: map[string]*dynamodb.Condition{
			"accountGuid": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("001F0000010Hrj0IAC"),
					},
				},
			},
		},
		ScanIndexForward: aws.Bool(false),
		Limit:            aws.Int64(2),
	}

	var resp1, err1 = dbSvc.Query(queryInput)

	if err1 != nil {
		fmt.Println("3333 Error: ", err1)
	} else {
		obj := []SSA{}
		err = dynamodbattribute.UnmarshalListOfMaps(resp1.Items, &obj)

		for _, ssa := range obj {
			fmt.Println("AccountGUID: ", ssa.AccountGUID)
			fmt.Println("SSAGUID: ", ssa.SsaGUID)

			if ssa.Aircraft != nil {
				if len(ssa.Aircraft.AircraftGUID) != 0 {
					fmt.Println("AircraftGUID: ", ssa.Aircraft.AircraftGUID)
				}
				fmt.Println("Tail: ", len(ssa.Aircraft.Tail))
			} else {
				fmt.Println("SSA Aircraft is Nil")
			}

			/*
					"aircraft": {
				    "make": "FALCON",
				    "serial": "8X-443",
				    "model": "8X"
				  },
			*/
		}
	}
}
