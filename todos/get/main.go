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
	Id      string `json:"id"`
	Title   string `json:"title"`
	Details string `json:"details"`
}

func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	sess := session.Must(session.NewSession(&aws.Config{
		Endpoint: aws.String(os.Getenv("DYNAMODB_ENDPOINT")),
	}))

	svc := dynamodb.New(sess)

	todoId := request.PathParameters["id"]

	result, err := svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(os.Getenv("DYNAMODB_TABLE")),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(todoId),
			},
		},
	})

	if err != nil {
		fmt.Println("Problem fetching TODO: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, nil
	}

	if len(result.Item) == 0 {
		return events.APIGatewayProxyResponse{StatusCode: 404}, nil
	}

	item := Item{}

	err = dynamodbattribute.UnmarshalMap(result.Item, &item)

	if err != nil {
		fmt.Println("Failed to UnmarshalMap item: ", err.Error())
		return events.APIGatewayProxyResponse{StatusCode: 500}, err
	}

	marshalledItem, err := json.Marshal(item)

	return events.APIGatewayProxyResponse{
		Body:       string(marshalledItem),
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
