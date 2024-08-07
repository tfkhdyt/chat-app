package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/chat-app/server/auth"
	"github.com/tfkhdyt/chat-app/server/user"
)

func RegisterRouter(app *fiber.App, userController *user.UserController, authController *auth.AuthController) {
	auth := app.Group("/auth")
	auth.Post("/register", userController.Create)
	auth.Post("/login", authController.Login)
}
