package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Getstatus(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}
