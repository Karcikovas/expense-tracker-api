package router

import (
	"expense-tracker-api/internal/http/hello"
)

func ProvideRoutes(hello *hello.RouteHandler) []RouteHandler {
	return []RouteHandler{hello}
}
