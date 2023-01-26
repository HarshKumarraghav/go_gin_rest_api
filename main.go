package main

import (
	"context"
	"gogin-restapi/api/routes"
	"gogin-restapi/pkg/configs"
	"gogin-restapi/pkg/data"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	app := gin.Default()
	godotenv.Load()
	config := configs.FromEnv()

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(config.MongoURI))
	if err != nil {
		log.Panic(err)
	}
	db := client.Database("gogin_restapi")
	dataRepo := data.NewRepo(db)
	routes.CreateRoutes(app, dataRepo.(*data.Repo))
	app.Run("localhost:" + os.Getenv("PORT"))
}
