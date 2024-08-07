package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tfkhdyt/fiber-toolbox/validation"
)

type UserController interface {
	Create(ctx *fiber.Ctx) error
}

type userController struct {
	service UserService
}

func NewUserController(service UserService) UserController {
	return &userController{service: service}
}

func (c *userController) Create(ctx *fiber.Ctx) error {
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
