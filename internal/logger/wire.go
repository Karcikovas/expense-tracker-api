package logger

import "github.com/google/wire"

var DepSet = wire.NewSet(
	NewLogger,
)
