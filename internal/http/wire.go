package http

import (
	"expense-tracker-api/internal/http/middleware"
	"github.com/google/wire"
)

var DepSet = wire.NewSet(
	middleware.DepSet,
	NewServer,
)
