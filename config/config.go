package config

import "go.uber.org/fx"

// Module config
var Module = fx.Options(
	fx.Provide(NewDBConnection,
		NewEchoHandler,
		NewRequestHandler,
		LoadEnvs),
)
