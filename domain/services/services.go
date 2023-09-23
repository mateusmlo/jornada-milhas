package service

import "go.uber.org/fx"

// Module exported services
var Module = fx.Options(
	fx.Provide(NewUserService),
	fx.Provide(NewAuthService),
)
