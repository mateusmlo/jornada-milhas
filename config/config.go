package config

import "go.uber.org/fx"

// Module config
var Module = fx.Options(
	fx.Provide(NewDBConnection),
	fx.Provide(GetLogger),
	fx.Provide(GetLogger().GetGormLogger),
	fx.Provide(GetLogger().GetGinLogger),
	fx.Provide(NewRequestHandler),
	fx.Provide(LoadEnvs),
)
