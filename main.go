package main

import (
	"LoveLetGoer/configs"
	"LoveLetGoer/routes"
	"fmt"
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

	// / endpoint
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "landing", gin.H{})
	})

	router.POST("/create", func(ctx *gin.Context) {
		name := ctx.PostForm("roomName")
		user := ctx.PostForm("userName")
		configs.DB[name] = []string{user}
		ctx.Header("HX-Redirect", fmt.Sprintf("/%s", name))
		ctx.Status(200)
	})

	router.POST("/join", func(ctx *gin.Context) {
		name := ctx.PostForm("roomName")
		user := ctx.PostForm("userName")
		_, ok := configs.DB[name]
		if ok {
			configs.DB[name] = append(configs.DB[name], user)
			ctx.HTML(http.StatusOK, "room", gin.H{"title": name})
			return
		} else {
			configs.DB[name] = []string{user}
		}
		ctx.Header("HX-Redirect", fmt.Sprintf("/%s", name))
		ctx.Status(200)
	})

	router.GET("/:name", routes.GetRoom)

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
