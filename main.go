package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gitlab.com/newsunbanjade/golang_mongo/controllers"
	"gitlab.com/newsunbanjade/golang_mongo/services"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	UserController controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func init() {
	ctx = context.TODO()
	mongoconn := options.Client().ApplyURI("<INSERT YOUR MONGO DB URL>")
	mongoclient, err = mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}
	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Mongo Is Connected")
	usercollection = mongoclient.Database("golang").Collection("users")
	userservice = services.NewUserService(usercollection, ctx)
	UserController = controllers.New(userservice)
	server = gin.Default()
}

func main() {
	defer mongoclient.Disconnect(ctx)
	basepath := server.Group("/v1")
	UserController.RegisterUserRoutes(basepath)

	log.Fatal(server.Run(":9091"))

}
