package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/fiber-toolbox/validation"
)

type AuthController struct {
	service AuthService
}

func NewAuthController(service AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) Login(ctx *fiber.Ctx) error {
	payload := new(LoginRequestDTO)
	if err := validation.ValidateBody(ctx, payload); err != nil {
		return err
	}

	resp, err := c.service.login(payload)
	if err != nil {
		return err
	}

	return ctx.JSON(resp)
}
