package repository

import "go.uber.org/fx"

// Module exports repositories
var Module = fx.Options(
	fx.Provide(NewUserRepository, NewReviewRepository),
)
