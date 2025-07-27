package main

import (
	"LoveLetGoer/configs"
	"LoveLetGoer/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	configs.ConfigureDB()

	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// status endpoint
	router.GET("/status", routes.Getstatus)

	// landing page endpoint
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "landing", gin.H{})
	})

	{
		room := router.Group("/room")
		room.GET("/:name", routes.GetRoom)
		room.POST("/create", routes.CreateRoom)
		room.POST("/join", routes.JoinRoom)
	}

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
