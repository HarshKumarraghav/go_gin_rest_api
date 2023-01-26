package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	router := gin.Default()

	router.Run("localhost:" + os.Getenv("PORT"))

}
