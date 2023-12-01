package middleware

import (
	"log/slog"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(log *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		for _, err := range ctx.Errors {
			log.Error("Error:", "URL", ctx.Request.URL, "method", ctx.Request.Method, "error", err.Error())
			ctx.JSON(-1, err)
		}
	}
}
