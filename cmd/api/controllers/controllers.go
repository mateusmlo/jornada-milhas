package controllers

import "go.uber.org/fx"

// Module exported controllers
var Module = fx.Options(
	fx.Provide(NewUserController),
	fx.Provide(NewJWTAuthController),
	fx.Provide(NewReviewController),
)
