package http

import (
	"errors"
	"example.com/m/v2/internal/http/middleware"
	"example.com/m/v2/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	logger      logger.Service
	middlewares []middleware.MiddleWare
}

func NewServer(
	logger logger.Service,
	middlewares []middleware.MiddleWare,
) *Server {
	return &Server{
		logger:      logger,
		middlewares: middlewares,
	}
}

func (s *Server) Start() {
	ginEngine := gin.New()

	ginEngine.UseRawPath = true
	ginEngine.UnescapePathValues = false

	s.registerMiddleware(ginEngine)

	httpServer := http.Server{
		Addr:              "8000",
		Handler:           ginEngine,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := httpServer.ListenAndServe(); errors.Is(err, http.ErrServerClosed) {
		s.logger.LogError(err.Error())
	}

	s.logger.LogInfo("HTTP server Started")

}

func (s *Server) registerMiddleware(r *gin.Engine) {
	for _, middleware := range s.middlewares {
		r.Use(middleware.MiddleWareFunc())
	}

	r.Use(gin.Logger())
}
