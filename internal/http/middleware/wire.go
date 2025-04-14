package middleware

import (
	"github.com/google/wire"
)

var DepSet = wire.NewSet(
	NewLog,
	ProvideMiddlewares,
)
