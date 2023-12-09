package utils

import "github.com/gin-gonic/gin"

func GetPk(ctx *gin.Context) string {
	pk := ctx.Param("pk")
	return pk
}
