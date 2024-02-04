package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/data"
)

type UsersHandler struct {
	Ctx *data.ZdkContext
}

func NewUsersHandler(ctx *data.ZdkContext) *UsersHandler {
	return &UsersHandler{Ctx: ctx}
}

func (handler *UsersHandler) List(ctx *fiber.Ctx) error {
	// request := new(models.DTOUserListRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return err
	// }

	return ctx.SendString("")
}

func (handler *UsersHandler) New(ctx *fiber.Ctx) error {
	// request := new(models.DTOUserNewRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return err
	// }

	return ctx.SendString("")
}

func (handler *UsersHandler) Update(ctx *fiber.Ctx) error {
	// request := new(models.DTOUserUpdateRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return err
	// }

	return ctx.SendString("")
}

func (handler *UsersHandler) Delete(ctx *fiber.Ctx) error {
	// request := new(models.DTOUserDeleteRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return err
	// }

	return ctx.SendString("")
}
