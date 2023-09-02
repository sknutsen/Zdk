package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/models"
	"github.com/sknutsen/Zdk/internal/storage"
)

type ShoppingListsHandler struct {
	Storage *storage.ShoppingListsStorage
}

func NewShoppingListsHandler(storage *storage.ShoppingListsStorage) *ShoppingListsHandler {
	return &ShoppingListsHandler{Storage: storage}
}

func (handler *ShoppingListsHandler) List(ctx *fiber.Ctx) error {
	request := new(models.ShoppingListDTOListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListsHandler) New(ctx *fiber.Ctx) error {
	request := new(models.ShoppingListDTONewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListsHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.ShoppingListDTOUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ShoppingListsHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.ShoppingListDTODeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
