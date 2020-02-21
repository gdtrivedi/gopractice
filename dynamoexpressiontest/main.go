package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type AssetLookup struct {
	AircraftSID   string `json:"aircraft_sid"`
	InsertionTime int64  `json:"insertion_time"`
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

	var aircraftSidKeyName = "aircraft_sid"

	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String("gogo-dash-ast-persist-prod-assetlookup"),
		KeyConditions: map[string]*dynamodb.Condition{
			aircraftSidKeyName: {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						S: aws.String("0012A000024ie1rQAA:a01F000000OhuLOIAZ"),
					},
				},
			},
		},
		ScanIndexForward: aws.Bool(false),
		Limit:            aws.Int64(1),
	}

	var resp1, err1 = dbSvc.Query(queryInput)

	if err1 != nil {
		fmt.Println("3333 Error: ", err1)
	} else {
		obj := []AssetLookup{}
		err = dynamodbattribute.UnmarshalListOfMaps(resp1.Items, &obj)

		for _, asset := range obj {
			fmt.Println("AircraftSID: ", asset.AircraftSID)
			fmt.Println("InsertionTime: ", asset.InsertionTime)
		}
	}

	// filt := expression.Name("aircraft_sid").Equal(expression.Value("0012A000024ie1rQAA:a01F000000OhuLOIAZ"))
	keyCond := expression.Key("aircraft_sid").Equal(expression.Value("0012A000024ie1rQAA:a01F000000OhuLOIAZ"))
	proj := expression.NamesList(
		expression.Name("aircraft_sid"),
		expression.Name("insertion_time"),
	)

	expr, err := expression.NewBuilder().WithKeyCondition(keyCond).WithProjection(proj).Build()
	// expr, err := expression.NewBuilder().WithKeyCondition(keyCond).Build()
	if err != nil {
		fmt.Println(err)
	}

	input := &dynamodb.QueryInput{
		KeyConditionExpression:    expr.KeyCondition(),
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String("gogo-dash-ast-persist-prod-assetlookup"),
		ScanIndexForward:          aws.Bool(false),
		Limit:                     aws.Int64(2),
	}

	resultExpr, err := dbSvc.Query(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case dynamodb.ErrCodeProvisionedThroughputExceededException:
				fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
			case dynamodb.ErrCodeResourceNotFoundException:
				fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
			case dynamodb.ErrCodeInternalServerError:
				fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	obj1 := []AssetLookup{}
	err = dynamodbattribute.UnmarshalListOfMaps(resultExpr.Items, &obj1)

	for _, asset := range obj1 {
		fmt.Println("AircraftSID: ", asset.AircraftSID)
		fmt.Println("InsertionTime: ", asset.InsertionTime)
	}

	fmt.Println(resultExpr)
}
