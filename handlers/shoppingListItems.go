package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/models"
)

type ShoppingListItemsHandler struct {
	Ctx         *data.ZdkContext
	UserHandler *UserHandler
}

func NewShoppingListItemsHandler(ctx *data.ZdkContext, userHandler *UserHandler) *ShoppingListItemsHandler {
	return &ShoppingListItemsHandler{Ctx: ctx, UserHandler: userHandler}
}

func (handler *ShoppingListItemsHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOShoppingListItemListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListItemsHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOShoppingListItemListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListItemsHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOShoppingListItemListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListItemsHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOShoppingListItemListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
