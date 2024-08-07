package fx

import (
	"github.com/tfkhdyt/chat-app/server/auth"
	"github.com/tfkhdyt/chat-app/server/lib/fiber"
	"github.com/tfkhdyt/chat-app/server/lib/gorm"
	"github.com/tfkhdyt/chat-app/server/user"
	"go.uber.org/fx"
)

func NewFx() *fx.App {
	return fx.New(
		fx.Provide(
			// db
			gorm.NewGormDB,

			// user
			fx.Annotate(user.NewUserRepository, fx.As(new(user.UserRepo))),
			fx.Annotate(user.NewUserService, fx.As(new(user.UserService))),
			user.NewUserController,

			// auth
			fx.Annotate(auth.NewAuthService, fx.As(new(auth.AuthService))),
			auth.NewAuthController,
		),
		fx.Invoke(fiber.NewFiberServer),
	)
}
