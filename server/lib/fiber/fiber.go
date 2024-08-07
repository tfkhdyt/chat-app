package fiber

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/tfkhdyt/chat-app/server/router"
	"github.com/tfkhdyt/chat-app/server/user"
	"github.com/tfkhdyt/fiber-toolbox/exception"
	"go.uber.org/fx"
)

func NewFiberServer(lc fx.Lifecycle, userController user.UserController) *fiber.App {
	app := fiber.New(fiber.Config{
		ErrorHandler: exception.NewErrorHandler(),
	})

	app.Use(logger.New())

	router.RegisterRouter(app, userController)

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := app.Listen(":" + os.Getenv("PORT")); err != nil {
					panic(err)
				}
			}()

			return nil
		},
		OnStop: func(_ context.Context) error {
			if err := app.Shutdown(); err != nil {
				return err
			}

			return nil
		},
	})

	return app
}
