package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/data"
)

type SettingsHandler struct {
	Ctx         *data.ZdkContext
	UserHandler *UserHandler
}

func NewSettingsHandler(ctx *data.ZdkContext, userHandler *UserHandler) *SettingsHandler {
	return &SettingsHandler{Ctx: ctx, UserHandler: userHandler}
}

func (handler *SettingsHandler) List(ctx *fiber.Ctx) error {
	// request := new(models.DTOSettingListRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return err
	// }

	return ctx.SendString("")
}

func (handler *SettingsHandler) New(ctx *fiber.Ctx) error {
	// request := new(models.DTOSettingNewRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return err
	// }

	return ctx.SendString("")
}

func (handler *SettingsHandler) Update(ctx *fiber.Ctx) error {
	// request := new(models.DTOSettingUpdateRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return err
	// }

	return ctx.SendString("")
}

func (handler *SettingsHandler) Delete(ctx *fiber.Ctx) error {
	// request := new(models.DTOSettingDeleteRequest)

	// if err := ctx.BodyParser(request); err != nil {
	// 	return err
	// }

	return ctx.SendString("")
}
