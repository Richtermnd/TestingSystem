package server

import (
	"context"
	"log/slog"

	"github.com/Richtermnd/TestingSystem/internal/service"
	"github.com/Richtermnd/TestingSystem/pkg/tests"
	"github.com/Richtermnd/TestingSystem/server/handler"
	"github.com/Richtermnd/TestingSystem/server/middleware"
	"github.com/gin-gonic/gin"
)

func Run(
	ctx context.Context,
	log *slog.Logger,
	addr string,
) error {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())
	engine.Use(middleware.ErrorHandler(log))

	apiRouter := engine.Group("/api")
	// CRUD for tests.
	handler.CRUDRouter[tests.Test](
		apiRouter,
		"/tests",
		service.NewTestService(log),
	)

	engine.Run(addr)
	return nil
}
