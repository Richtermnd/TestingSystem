package handler

import (
	"log/slog"

	"github.com/Richtermnd/TestingSystem/server/handlers/api"
	"github.com/Richtermnd/TestingSystem/server/handlers/global"
	"github.com/gin-gonic/gin"
)

// Server Adapter to avoid cycle import. It is cursed?
type Server interface {
	GetApiRouter() *gin.RouterGroup
	GetGlobalRouter() *gin.RouterGroup
	GetLogger() *slog.Logger
}

func Register(server Server) {
	log := server.GetLogger()
	apiRouter := server.GetApiRouter()
	globalRouter := server.GetGlobalRouter()
	api.RegisterTestHandlers(apiRouter, log)
	global.Register(globalRouter, log)
}
