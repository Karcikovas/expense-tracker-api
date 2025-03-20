package http

import (
	//"expense-tracker-api/internal/http/middleware"
	"expense-tracker-api/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	logger logger.Service
	//middlewares []middleware.MiddleWare
}

func NewServer(
	logger logger.Service,
	// middlewares []middleware.MiddleWare,
) *Server {
	return &Server{
		logger: logger,
		//middlewares: middlewares,
	}
}

func (s *Server) Start() {
	ginEngine := gin.New()

	ginEngine.UseRawPath = true
	ginEngine.UnescapePathValues = false

	//s.registerMiddleware(ginEngine)

	httpServer := http.Server{
		Addr:              "localhost:8000",
		Handler:           ginEngine,
		ReadHeaderTimeout: 10 * time.Second,
	}

	ginEngine.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	})

	err := httpServer.ListenAndServe()

	if err != nil {
		s.logger.LogError(err.Error())
	}

	s.logger.LogInfo("HTTP server Started")

}

func (s *Server) registerMiddleware(r *gin.Engine) {
	//for _, middleware := range s.middlewares {
	//	r.Use(middleware.MiddleWareFunc())
	//}
	//
	//r.Use(gin.Logger())
}
