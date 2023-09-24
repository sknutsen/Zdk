package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/models"
	"github.com/sknutsen/Zdk/internal/storage"
)

type SettingsHandler struct {
	Storage *storage.ShoppingListItemsStorage
}

func NewSettingsHandler(storage *storage.ShoppingListItemsStorage) *SettingsHandler {
	return &SettingsHandler{Storage: storage}
}

func (handler *SettingsHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOSettingListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *SettingsHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOSettingNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *SettingsHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOSettingUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *SettingsHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOSettingDeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
