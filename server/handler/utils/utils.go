package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func IdFromParams(ctx *gin.Context) (*primitive.ObjectID, error) {
	param_id := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(param_id)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return nil, err
	}
	return &id, nil
}
