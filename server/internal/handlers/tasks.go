package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/models"
	"github.com/sknutsen/Zdk/internal/storage"
)

type TasksHandler struct {
	Storage *storage.ShoppingListItemsStorage
}

func NewTasksHandler(storage *storage.ShoppingListItemsStorage) *TasksHandler {
	return &TasksHandler{Storage: storage}
}

func (handler *TasksHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *TasksHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	handler.Storage.ZdkCtx.DB.Create(models.Task{
		Name:   request.Name,
		UserId: "",
	})

	return ctx.SendString("")
}

func (handler *TasksHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *TasksHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskDeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
