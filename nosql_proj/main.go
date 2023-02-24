package main

import (
	"github.com/DigitalTroller/nosql_proj/configs"
	"github.com/DigitalTroller/nosql_proj/middleware"
	"github.com/DigitalTroller/nosql_proj/routes" //add this
	"github.com/gin-gonic/gin"
	_ "github.com/heroku/x/hmetrics/onload"
	"net/http"
)

func main() {
	r := gin.New()

	configs.ConnectDB()
	r.Static("/css", "./css")

	r.Use(gin.Logger())
	routes.UserRoutes(r)

	r.LoadHTMLGlob("web_pages/*")

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main website",
		})
	})
	r.GET("/:filename", func(c *gin.Context) {
		c.HTML(http.StatusOK, c.Param("filename")+".html", gin.H{})
	})

	r.Use(middleware.Authentication())

	// API-2
	r.GET("/api-1", func(c *gin.Context) {

		c.JSON(200, gin.H{"success": "Access granted for api-1"})

	})

	// API-1
	r.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Access granted for api-2"})
	})

	r.Run(":" + "4000")

}
