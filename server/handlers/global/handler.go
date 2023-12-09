package global

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "pong")
}

func Register(router *gin.RouterGroup, log *slog.Logger) {
	router.GET("/ping", ping)
}
