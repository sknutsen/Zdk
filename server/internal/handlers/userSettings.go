package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/models"
	"github.com/sknutsen/Zdk/internal/storage"
)

type UserSettingsHandler struct {
	Storage *storage.UsersStorage
}

func NewUserSettingsHandler(storage *storage.UsersStorage) *UserSettingsHandler {
	return &UserSettingsHandler{Storage: storage}
}

func (handler *UserSettingsHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOUserSettingListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *UserSettingsHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOUserSettingNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *UserSettingsHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOUserSettingUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *UserSettingsHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOUserSettingDeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}
