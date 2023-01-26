package main

import (
	"context"
	"gogin-restapi/api/routes"
	"gogin-restapi/pkg/configs"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	godotenv.Load()
	router := gin.Default()
	config := configs.FromEnv()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Panic(err)
	}
	// db := client.Database("gogin_restapi")
	routes.CreateRoutes(router)
	router.Run("localhost:" + os.Getenv("PORT"))

}
