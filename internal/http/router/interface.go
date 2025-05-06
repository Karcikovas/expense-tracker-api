package router

import "github.com/gin-gonic/gin"

type RouteHandler interface {
	HttpHandler() (string, string, gin.HandlerFunc)
}
