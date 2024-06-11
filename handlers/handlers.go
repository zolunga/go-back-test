package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/zolunga/go-back-test/models"
)

func Handlers(ctx context.Context, req events.APIGatewayProxyRequest) models.Response {
	fmt.Printf("Processing path %s > %s", ctx.Value(models.Key("path")).(string), ctx.Value(models.Key("method")).(string))
	var r models.Response
	r.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
	}

	r.Message = "Method invalid"
	return r
}
