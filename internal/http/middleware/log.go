package middleware

import (
	"expense-tracker-api/internal/logger"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Log struct {
	logger logger.Service
}

func NewLog(logger logger.Service) *Log {
	return &Log{
		logger: logger,
	}
}

func (l *Log) MiddleWareFunc() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			l.logger.LogInfo(fmt.Sprintf(`IP: %s`, ctx.RemoteIP()))
			l.logger.LogInfo(fmt.Sprintf(`Request method: %s %s`, ctx.Request.Method, ctx.Request.RequestURI))
		}()

		ctx.Next()
	}
}
