package routes

import (
	"LoveLetGoer/configs"
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
