package api

import (
	"net/http"

	"github.com/Richtermnd/TestingSystem/server/utils"
	"github.com/gin-gonic/gin"
)

// CRUDService Service that can make classic CRUD operations
type CRUDService[T any] interface {
	Create(T) (string, error)
	Read(pk string) (*T, error)
	ReadAll() ([]*T, error)
	Update(pk string, updatedItem *T) (string, error)
	Delete(pk string) (bool, error)
}

// CRUDRouter create router and add CRUD handlers if Service implement CRUDService interface.
func CRUDRouter[T any](
	apiRouter *gin.RouterGroup,
	prefix string,
	service CRUDService[T],
) *gin.RouterGroup {
	router := apiRouter.Group(prefix)
	// Create
	router.POST("/", func(ctx *gin.Context) {
		var item T
		if err := ctx.Bind(&item); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		pk, err := service.Create(item)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, pk)
	})

	// ReadAll
	router.GET("/", func(ctx *gin.Context) {
		items, err := service.ReadAll()
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, &items)
	})

	// Read
	router.GET("/:pk", func(ctx *gin.Context) {
		pk := utils.GetPk(ctx)
		item, err := service.Read(pk)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		ctx.JSON(http.StatusOK, *item)
	})

	// Update
	router.PUT("/:pk", func(ctx *gin.Context) {
		// Getting item from body
		var item T
		if err := ctx.Bind(&item); err != nil {
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		// Getting pk from params
		pk := utils.GetPk(ctx)
		// Updating item.
		service.Update(pk, &item)
	})

	// Delete
	router.DELETE("/:pk", func(ctx *gin.Context) {
		pk := utils.GetPk(ctx)
		res, err := service.Delete(pk)
		if err != nil {
			ctx.AbortWithError(http.StatusInternalServerError, err)
		}
		ctx.JSON(http.StatusOK, res)
	})

	return router
}
