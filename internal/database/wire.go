package database

import "github.com/google/wire"

var DepSet = wire.NewSet(
	NewDBCon,
)
