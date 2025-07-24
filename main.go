package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	// status endpoint
	router.GET("/status", func(c *gin.Context) {
		c.String(http.StatusOK, "OK")
	})

	// / endpoint
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "landing", gin.H{})
	})

	// Serve static files (CSS, JS, etc.)
	// r.Static("/static", "./static")

	// Routes
	// r.GET("/", homeHandler)

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
