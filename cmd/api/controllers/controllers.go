package controllers

import "go.uber.org/fx"

// Module exported controllers
var Module = fx.Options(
	fx.Provide(NewUserController, NewAuthController, NewReviewController),
)
