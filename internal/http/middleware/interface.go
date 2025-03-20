package middleware

import "github.com/gin-gonic/gin"

type MiddleWare interface {
	MiddleWareFunc() gin.HandlerFunc
}
