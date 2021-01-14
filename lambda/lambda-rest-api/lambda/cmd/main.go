package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"lambda-rest-api/pkg/dynamoHandlers"
	"net/http"
	"os"
)

var (
	dynamoDbClient dynamodbiface.DynamoDBAPI
)

const tableName = "lambdaingouser"

func main() {
	region := os.Getenv("AWS_REGION")

	awsSession, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		return
	}

	dynamoDbClient = dynamodb.New(awsSession)
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Printf("table name is %s", tableName)

	switch request.HTTPMethod {
	case http.MethodGet:
		return dynamoHandlers.GetUser(request, tableName, dynamoDbClient)
	case http.MethodPut:
		return dynamoHandlers.UpdateUser(request, tableName, dynamoDbClient)
	case http.MethodPost:
		return dynamoHandlers.CreateUser(request, tableName, dynamoDbClient)
	case http.MethodDelete:
		return dynamoHandlers.DeleteUser(request, tableName, dynamoDbClient)
	default:
		fmt.Printf("method type is %s", request.HTTPMethod)
		return dynamoHandlers.UnhandledMethod(request.Body, request.HTTPMethod)
	}
}
