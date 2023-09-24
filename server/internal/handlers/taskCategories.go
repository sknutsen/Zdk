package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/models"
	"github.com/sknutsen/Zdk/internal/storage"
)

type TaskCategoriesHandler struct {
	Storage *storage.ShoppingListItemsStorage
}

func NewTaskCategoriesHandler(storage *storage.ShoppingListItemsStorage) *TaskCategoriesHandler {
	return &TaskCategoriesHandler{Storage: storage}
}

func (handler *TaskCategoriesHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskCategoryListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *TaskCategoriesHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskCategoryNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *TaskCategoriesHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskCategoryUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *TaskCategoriesHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskCategoryDeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
