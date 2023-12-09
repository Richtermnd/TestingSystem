package api

import (
	"log/slog"
	"net/http"

	"github.com/Richtermnd/TestingSystem/server/utils"
	"github.com/Richtermnd/TestingSystem/service"
	tests "github.com/Richtermnd/TestingSystem/testingSystem"
	"github.com/gin-gonic/gin"
)

func checkTest(log *slog.Logger, service *service.TestService) func(*gin.Context) {
	return func(ctx *gin.Context) {
		pk := utils.GetPk(ctx)
		log.Info("Check test", "pk", pk)
		test, err := service.Read(pk)
		if err != nil {
			log.Error(err.Error())
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		var answers *tests.QuestionAnswers
		if err := ctx.Bind(&answers); err != nil {
			log.Error(err.Error())
			ctx.AbortWithError(http.StatusBadRequest, err)
			return
		}
		res := answers.Check(test)
		log.Info("Checking test result", "pk", pk, "res", res)
		ctx.JSON(http.StatusOK, &TestCheckResponse{IsPassed: res})
	}
}

func RegisterTestHandlers(
	router *gin.RouterGroup,
	log *slog.Logger,
) {
	service := service.NewTestService(log)
	r := CRUDRouter[tests.Test](router, "/tests", service)
	r.POST("/:pk/check", checkTest(log, service))
}
