package routes

import (
	"LoveLetGoer/configs"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetRoom(ctx *gin.Context) {
	var name string = ctx.Param("name")
	_, ok := configs.DB[name]
	if ok {
		ctx.HTML(http.StatusOK, "room", gin.H{"title": name})
	} else {
		ctx.String(http.StatusNotFound, "")
	}
}

func CreateRoom(ctx *gin.Context) {
	name := ctx.PostForm("roomName")
	user := ctx.PostForm("userName")
	configs.DB[name] = []string{user}
	ctx.Header("HX-Redirect", fmt.Sprintf("room/%s", name))
	ctx.Status(200)
}

func JoinRoom(ctx *gin.Context) {
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
	ctx.Header("HX-Redirect", fmt.Sprintf("room/%s", name))
	ctx.Status(200)
}
