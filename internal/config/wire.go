package config

import "github.com/google/wire"

var DepSet = wire.NewSet(
	NewConfig,
)
