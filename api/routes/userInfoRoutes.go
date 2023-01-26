package routes

import (
	"gogin-restapi/pkg/data"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUserHandler(repo *data.Repo) gin.HandlerFunc {
	return func(c *gin.Context) {
		var user data.InUser

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError, "Message": "error", "Data": map[string]interface{}{"data": err.Error()}})
			return

		}

		response, err := repo.Create(user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"Status": http.StatusInternalServerError, "Message": "error", "Data": map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"user_id": response.ID, "status": "success"})
		return
	}
}
func CreateRoutes(app *gin.Engine, userRepo *data.Repo) {
	app.GET("/createuser", CreateUserHandler(userRepo))
}
