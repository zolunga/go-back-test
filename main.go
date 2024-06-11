package main

import (
	"context"
	"os"
	"strings"

	events "github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/zolunga/go-back-test/awsgo"
	"github.com/zolunga/go-back-test/bd"
	"github.com/zolunga/go-back-test/handlers"
	"github.com/zolunga/go-back-test/models"
	"github.com/zolunga/go-back-test/secretmanager"
)

func main() {
	lambda.Start(RunLambda)
}

func RunLambda(ctx context.Context, req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse
	awsgo.InitAWS()
	if !ValidateEnv() {
		res = createError(400, "Bad config at ENV")
		return res, nil
	}
	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))
	if err != nil {
		res = createError(400, "Bad config at secrets")
		return res, nil
	}
	path := strings.Replace(req.PathParameters["twitter"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), req.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("mongouri"), SecretModel.Mongouri)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.Jwtsign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), req.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// database connection
	err = bd.ConnectMongo(awsgo.Ctx)
	if err != nil {
		res = createError(500, "Error at mongo db connection")
		return res, nil
	}

	response := handlers.Handlers(awsgo.Ctx, req)
	if response.CustomResponse == nil {
		return createError(response.Status, response.Message), nil
	} else {
		return response.CustomResponse, nil
	}

}

func ValidateEnv() bool {
	_, getParam := os.LookupEnv("SecretName")
	if !getParam {
		return getParam
	}

	_, getParam = os.LookupEnv("BucketName")
	if !getParam {
		return getParam
	}

	_, getParam = os.LookupEnv("UrlPrefix")
	if !getParam {
		return getParam
	}

	return getParam
}

func createError(code int, msg string) *events.APIGatewayProxyResponse {
	return &events.APIGatewayProxyResponse{
		StatusCode: code,
		Body:       msg,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}
}
