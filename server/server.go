package server

import (
	"log/slog"

	"github.com/Richtermnd/TestingSystem/server/handlers"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine       *gin.Engine
	log          *slog.Logger
	apiRouter    *gin.RouterGroup // Api router for api requests.
	globalRouter *gin.RouterGroup // global router for pages and other.
}

func (s *Server) Init() {
	// Adding middlewares
	s.engine.Use(gin.Recovery())
	s.engine.Use(gin.Logger()) // Мб написать свой?

	// Register handlers
	handler.Register(s)
}

func (s *Server) Run() {
	s.engine.Run(":8080")
}

func NewServer(log *slog.Logger) *Server {
	engine := gin.New()
	apiRouter := engine.Group("/api")
	globalRouter := engine.Group("/")
	return &Server{
		engine:       engine,
		log:          log,
		apiRouter:    apiRouter,
		globalRouter: globalRouter,
	}
}

func (s *Server) GetApiRouter() *gin.RouterGroup {
	return s.apiRouter
}

func (s *Server) GetGlobalRouter() *gin.RouterGroup {
	return s.globalRouter
}

func (s *Server) GetLogger() *slog.Logger {
	return s.log
}
