package http

import (
	"expense-tracker-api/internal/config"
	"log"

	//"expense-tracker-api/internal/http/middleware"
	"expense-tracker-api/internal/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Server struct {
	logger logger.Service
	config config.Config
	//middlewares []middleware.MiddleWare
}

func NewServer(
	logger logger.Service,
	config config.Config,
	// middlewares []middleware.MiddleWare,
) *Server {
	return &Server{
		logger: logger,
		config: config,
		//middlewares: middlewares,
	}
}

func (s *Server) Start() {
	ginEngine := gin.New()

	ginEngine.UseRawPath = true
	ginEngine.UnescapePathValues = false

	log.Println("localhost:" + s.config.HttpPort)

	//s.registerMiddleware(ginEngine)

	httpServer := http.Server{
		Addr:              "localhost:" + s.config.HttpPort,
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
