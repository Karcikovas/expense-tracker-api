package router

import "github.com/gin-gonic/gin"

type Route struct {
	HttpMethod   string
	RelativePath string
	Handlers     []gin.HandlerFunc
}

type Handler interface {
	HttpHandler() Route
}
