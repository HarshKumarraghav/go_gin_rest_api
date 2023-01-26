package routes

import "github.com/gin-gonic/gin"

func CreateUserHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
func CreateRoutes(app *gin.Engine) {
	app.GET("/createuser", CreateUserHandler())
}
