package bd

import (
	"context"
	"fmt"

	"github.com/zolunga/go-back-test/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCon *mongo.Client
var DatabaseName string

func ConnectMongo(ctx context.Context) error {
	// user := ctx.Value(models.Key("user")).(string)
	// pass := ctx.Value(models.Key("password")).(string)
	uri := ctx.Value(models.Key("mongouri")).(string)
	completeUri := fmt.Sprintf("%s?retryWrites=true&w=majority", uri)

	var clientOptions = options.Client().ApplyURI(completeUri)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Succesfully connected to mongo")
	MongoCon = client
	DatabaseName = ctx.Value(models.Key("database")).(string)
	return nil
}

func TestConnection() bool {
	err := MongoCon.Ping(context.TODO(), nil)
	return err == nil
}
