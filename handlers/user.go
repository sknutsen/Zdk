package handlers

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/models"
)

type UserHandler struct {
	Ctx            *data.ZdkContext
	SessionManager *auth.SessionManager
}

func NewUserHandler(ctx *data.ZdkContext, session *auth.SessionManager) *UserHandler {
	return &UserHandler{Ctx: ctx, SessionManager: session}
}

func (handler *UserHandler) GetUserProfile(ctx *fiber.Ctx) (*models.UserProfile, error) {
	session, err := handler.SessionManager.SessionStore.Get(ctx)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}

	profile := session.Get("profile")
	if profile == nil {
		return nil, errors.New("failed getting profile")
	}

	return getUserProfile(profile.(map[string]interface{})), nil
}

func (handler *UserHandler) GetUserId(ctx *fiber.Ctx) (string, error) {
	profile, err := handler.GetUserProfile(ctx)
	return profile.UserId, err
}

func getUserProfile(profile map[string]interface{}) *models.UserProfile {
	return &models.UserProfile{
		UserId: profile["sub"].(string),
	}
}
