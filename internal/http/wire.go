package http

import (
	"expense-tracker-api/internal/http/hello"
	"expense-tracker-api/internal/http/middleware"
	"expense-tracker-api/internal/http/router"
	"github.com/google/wire"
)

var DepSet = wire.NewSet(
	hello.NewRouteHandler,
	middleware.DepSet,
	router.DepSet,
	NewServer,
)
