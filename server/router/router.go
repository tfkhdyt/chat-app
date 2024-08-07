package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/chat-app/server/user"
)

func RegisterRouter(app *fiber.App, userController user.UserController) {
	auth := app.Group("/auth")
	auth.Post("/register", userController.Create)
}
