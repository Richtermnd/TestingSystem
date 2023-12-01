package handler

import (
	"net/http"

	"github.com/Richtermnd/TestingSystem/internal/service"
	"github.com/Richtermnd/TestingSystem/server/handler/utils"
	"github.com/gin-gonic/gin"
)

func CRUDRouter[T any](
	apiRouter *gin.RouterGroup,
	prefix string,
	service service.IService[T],
) *gin.RouterGroup {
	router := apiRouter.Group(prefix)

	// Create
	router.POST("/", func(ctx *gin.Context) {
		var item T
		if err := ctx.Bind(&item); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		id, err := service.Create(item)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, id)
	})

	// ReadAll
	router.GET("/", func(ctx *gin.Context) {
		items, err := service.ReadAll()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, items)
	})

	// ReadOne
	router.GET("/:id", func(ctx *gin.Context) {
		id, err := utils.IdFromParams(ctx)
		if err != nil {
			return
		}
		item, err := service.ReadOne(id)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, item)
	})

	// Update
	router.PUT("/:id", func(ctx *gin.Context) {
		// Getting item from body
		var item T
		if err := ctx.Bind(&item); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		// Getting id from params
		id, err := utils.IdFromParams(ctx)
		if err != nil {
			return
		}
		// Updating item.
		service.Update(id, item)
	})

	// Delete
	router.DELETE("/:id", func(ctx *gin.Context) {
		id, err := utils.IdFromParams(ctx)
		if err != nil {
			return
		}
		res, err := service.Delete(id)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, res)
	})

	return router
}
