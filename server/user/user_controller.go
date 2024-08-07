package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/fiber-toolbox/validation"
)

type UserController struct {
	service UserService
}

func NewUserController(service UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) Create(ctx *fiber.Ctx) error {
	user := new(UserRegisterDTO)
	if err := validation.ValidateBody(ctx, user); err != nil {
		return err
	}

	response, err := c.service.create(user)
	if err != nil {
		return err
	}

	return ctx.JSON(response)
}
