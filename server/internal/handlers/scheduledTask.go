package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/models"
	"github.com/sknutsen/Zdk/internal/storage"
)

type ScheduledTasksHandler struct {
	Storage *storage.ShoppingListItemsStorage
}

func NewScheduledTasksHandler(storage *storage.ShoppingListItemsStorage) *ScheduledTasksHandler {
	return &ScheduledTasksHandler{Storage: storage}
}

func (handler *ScheduledTasksHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOScheduledTaskListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ScheduledTasksHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOScheduledTaskNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ScheduledTasksHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOScheduledTaskUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *ScheduledTasksHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOScheduledTaskDeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
