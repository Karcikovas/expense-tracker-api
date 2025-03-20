package http

import "github.com/google/wire"

var DepSet = wire.NewSet(
	NewServer,
)
