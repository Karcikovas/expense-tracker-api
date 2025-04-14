package middleware

import (
	"expense-tracker-api/internal/logger"
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
			l.logger.LogInfo(ctx.Request.RequestURI)
		}()

		ctx.Next()
	}
}
