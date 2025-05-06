package http

import (
	"expense-tracker-api/internal/config"
	"expense-tracker-api/internal/http/middleware"
	"expense-tracker-api/internal/http/router"
	"log"

	"expense-tracker-api/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	logger      logger.Service
	config      config.Config
	middlewares []middleware.MiddleWare
	routes      []router.RouteHandler
}

func NewServer(
	logger logger.Service,
	config config.Config,
	middlewares []middleware.MiddleWare,
	routes []router.RouteHandler,
) *Server {
	return &Server{
		logger:      logger,
		config:      config,
		middlewares: middlewares,
		routes:      routes,
	}
}

func (s *Server) Start() {
	ginEngine := gin.New()

	ginEngine.UseRawPath = true
	ginEngine.UnescapePathValues = false

	log.Println("localhost:" + s.config.HttpPort)

	s.registerMiddleware(ginEngine)

	httpServer := http.Server{
		Addr:              "localhost:" + s.config.HttpPort,
		Handler:           ginEngine,
		ReadHeaderTimeout: 10 * time.Second,
	}

	s.registerRouter(ginEngine)

	err := httpServer.ListenAndServe()

	if err != nil {
		s.logger.LogError(err.Error())
	}

	s.logger.LogInfo("HTTP server Started")

}

func (s *Server) registerRouter(r *gin.Engine) {
	for _, route := range s.routes {
		r.Handle(route.HttpHandler())
	}
}

func (s *Server) registerMiddleware(r *gin.Engine) {
	for _, middleware := range s.middlewares {
		r.Use(middleware.MiddleWareFunc())
	}

	r.Use(gin.Logger())
}
