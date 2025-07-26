package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	db := make(map[string][]string)

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

	router.POST("/create", func(ctx *gin.Context) {
		name := ctx.PostForm("roomName")
		user := ctx.PostForm("userName")
		db[name] = []string{user}
		ctx.Header("HX-Redirect", fmt.Sprintf("/%s", name))
		ctx.Status(200)
	})

	router.POST("/join", func(ctx *gin.Context) {
		name := ctx.PostForm("roomName")
		user := ctx.PostForm("userName")
		_, ok := db[name]
		if ok {
			db[name] = append(db[name], user)
			ctx.HTML(http.StatusOK, "room", gin.H{"title": name})
			return
		} else {
			db[name] = []string{user}
		}
		ctx.Header("HX-Redirect", fmt.Sprintf("/%s", name))
		ctx.Status(200)
	})

	router.GET("/:name", func(ctx *gin.Context) {
		var name string = ctx.Param("name")
		_, ok := db[name]
		if ok {
			ctx.HTML(http.StatusOK, "room", gin.H{"title": name})
			return
		}
		ctx.Redirect(http.StatusFound, "/")
	})

	return router
}

func main() {
	router := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	router.Run(":8080")
}
