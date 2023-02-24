package routes

import (
	"github.com/DigitalTroller/nosql_proj/controllers"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/task_exec/:id", controllers.TaskExec())
	incomingRoutes.POST("/users/signup", controllers.SignUp())
	incomingRoutes.POST("/users/login", controllers.Login())
	incomingRoutes.POST("/file-upload", controllers.FileUpload())
	incomingRoutes.GET("/api-1", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/task")

	})
	incomingRoutes.POST("/addNewTask", controllers.AddNewTask())
	//incomingRoutes.GET("/", controllers.Main())
	incomingRoutes.LoadHTMLGlob("web_pages/*")

	//incomingRoutes.GET("/tasks.html", controllers.GetTasks())

	incomingRoutes.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	incomingRoutes.GET("/tasks", controllers.GetTasks())
	incomingRoutes.GET("/:filename", func(c *gin.Context) {
		c.HTML(http.StatusOK, c.Param("filename")+".html", gin.H{})
	})
}
