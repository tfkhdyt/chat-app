package lib

import (
	"github.com/tfkhdyt/chat-app/server/user"
	"go.uber.org/fx"
)

func NewFx() *fx.App {
	return fx.New(
		fx.Provide(
			// db
			NewGormDB,

			// user
			fx.Annotate(user.NewUserRepository, fx.As(new(user.UserRepo))),
			fx.Annotate(user.NewUserService, fx.As(new(user.UserService))),
			fx.Annotate(user.NewUserController, fx.As(new(user.UserController))),
		),
		fx.Invoke(NewFiberServer),
	)
}
