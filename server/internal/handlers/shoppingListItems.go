package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/models"
	"github.com/sknutsen/Zdk/internal/storage"
)

type ShoppingListItemsHandler struct {
	Storage *storage.ShoppingListItemsStorage
}

func NewShoppingListItemsHandler(storage *storage.ShoppingListItemsStorage) *ShoppingListItemsHandler {
	return &ShoppingListItemsHandler{Storage: storage}
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
