package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/models"
)

type ShoppingListsHandler struct {
	Ctx         *data.ZdkContext
	UserHandler *UserHandler
}

func NewShoppingListsHandler(ctx *data.ZdkContext, userHandler *UserHandler) *ShoppingListsHandler {
	return &ShoppingListsHandler{Ctx: ctx, UserHandler: userHandler}
}

func (handler *ShoppingListsHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOShoppingListNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListsHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOShoppingListNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListsHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOShoppingListUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListsHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOShoppingListDeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
