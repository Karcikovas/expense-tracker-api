package cmd

import "github.com/google/wire"

var DepSet = wire.NewSet(
	NewServeCMD,
	ProvideCMD,
	NewRoot,
)
