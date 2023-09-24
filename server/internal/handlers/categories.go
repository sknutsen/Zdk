package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/models"
	"github.com/sknutsen/Zdk/internal/storage"
)

type CategoriesHandler struct {
	Storage *storage.ShoppingListItemsStorage
}

func NewCategoriesHandler(storage *storage.ShoppingListItemsStorage) *CategoriesHandler {
	return &CategoriesHandler{Storage: storage}
}

func (handler *CategoriesHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOCategoryListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *CategoriesHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOCategoryNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *CategoriesHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOCategoryUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *CategoriesHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOCategoryDeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
