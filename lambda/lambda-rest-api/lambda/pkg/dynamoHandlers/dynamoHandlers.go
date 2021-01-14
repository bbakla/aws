package dynamoHandlers

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"lambda-rest-api/pkg/user"
	"net/http"
)

type ErrorBody struct {
	ErrorMessage *string `json:"error,omitempty"`
}

func UnhandledMethod(body string, method string) (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, fmt.Sprintf("method %s not allowed.\n Body is %s", method, body))
}

func GetUser(request events.APIGatewayProxyRequest, tableName string, dynamodbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {

	email := request.QueryStringParameters["email"]
	var result *user.User
	var err error
	if len(email) == 0 {
		return GetUsers(tableName, dynamodbClient)
	}

	result, err = user.FetchUser(email, tableName, dynamodbClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	return apiResponse(http.StatusOK, result)
}

func GetUsers(tableName string, dynamodbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.FetchUsers(tableName, dynamodbClient)

	status := http.StatusBadRequest
	var responseBody interface{}
	responseBody = result

	if err != nil {
		responseBody = ErrorBody{aws.String(err.Error())}
	}

	return apiResponse(status, responseBody)
}
func CreateUser(request events.APIGatewayProxyRequest, tableName string, dynamodbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(request, tableName, dynamodbClient)

	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	return apiResponse(http.StatusCreated, result)
}
func UpdateUser(request events.APIGatewayProxyRequest, tableName string, dynamodbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(request, tableName, dynamodbClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}
func DeleteUser(request events.APIGatewayProxyRequest, tableName string, dynamodbClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	err := user.DeleteUser(request, tableName, dynamodbClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusNoContent, nil)
}
