package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/models"
)

type UserSettingsHandler struct {
	Ctx         *data.ZdkContext
	UserHandler *UserHandler
}

func NewUserSettingsHandler(ctx *data.ZdkContext, userHandler *UserHandler) *UserSettingsHandler {
	return &UserSettingsHandler{Ctx: ctx, UserHandler: userHandler}
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

	userId, err := handler.UserHandler.GetUserId(ctx)
	if err != nil {
		return err
	}

	object := models.UserSetting{
		UserSettingId: uuid.NewString(),
		UserId:        userId,
		SettingId:     request.SettingId,
		Value:         request.Value,
		DateCreated:   time.Now(),
		DateChanged:   time.Now(),
	}

	handler.Ctx.DB.Create(object)

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
