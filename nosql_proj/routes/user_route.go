package routes

import (
	"github.com/DigitalTroller/nosql_proj/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	//incomingRoutes.GET("/", controllers.Main())
}
