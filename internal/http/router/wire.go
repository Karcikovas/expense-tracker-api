package router

import (
	"github.com/google/wire"
)

var DepSet = wire.NewSet(
	ProvideRoutes,
)
