package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Item struct {
	Id      string `json:"id,omitempty"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(os.Getenv("DYNAMODB_ENDPOINT")),
		//Credentials: credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), os.Getenv("AWS_ACCESS_KEY_TOKEN")), //injected to serverless.yml
	}))

	svc := dynamodb.New(sess)

	params := &dynamodb.ScanInput{
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
	}

	// scan the table
	result, err := svc.Scan(params)

	if err != nil {
		fmt.Println("Query API failed: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	var itemArray []Item

	for _, i := range result.Items {
		item := Item{}

		// result is of type *dynamodb.GetItemOutput
		// result.Item is of type map[string]*dynamodb.AttributeValue
		// UnmarshallMap result.item to item
		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			fmt.Println("Error unmarshaling: ", err.Error())
			return events.APIGatewayProxyResponse{StatusCode: 500}, nil
		}

		itemArray = append(itemArray, item)
	}

	fmt.Println("itemArray: ", itemArray)

	itemArrayString, err := json.Marshal(itemArray)
	if err != nil {
		fmt.Println("Error marshaling to JSON: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	return events.APIGatewayProxyResponse{Body: string(itemArrayString), StatusCode: 200}, nil
}

func main() {
	lambda.Start(Handler)
}
