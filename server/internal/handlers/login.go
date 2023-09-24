package handlers

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/auth"
	"github.com/sknutsen/Zdk/internal/config"
	"go.uber.org/zap"
)

type LoginHandler struct {
	log            *zap.Logger
	Authenticator  *auth.Authenticator
	SessionManager *auth.SessionManager
	Config         *config.Config
}

func NewLoginHandler(log *zap.Logger, authenticator *auth.Authenticator, sessionManager *auth.SessionManager, config *config.Config) *LoginHandler {
	return &LoginHandler{log: log, Authenticator: authenticator, SessionManager: sessionManager, Config: config}
}

func (handler *LoginHandler) Login(ctx *fiber.Ctx) error {
	state, err := generateRandomState()
	if err != nil {
		return err
	}

	handler.log.Info("callback: state ", zap.String("state", state))
	session, _ := handler.SessionManager.SessionStore.Get(ctx)
	session.Set("state", state)
	if err := session.Save(); err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.Redirect(handler.Authenticator.AuthCodeURL(state), http.StatusTemporaryRedirect)
}

func generateRandomState() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	state := base64.StdEncoding.EncodeToString(b)

	return state, nil
}
