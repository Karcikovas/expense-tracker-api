package hello

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type RouteHandler struct {
}

func NewRouteHandler() *RouteHandler {
	return &RouteHandler{}
}

func (h *RouteHandler) HttpHandler() (string, string, gin.HandlerFunc) {
	return "GET", "/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome Gin Server")
	}
}
