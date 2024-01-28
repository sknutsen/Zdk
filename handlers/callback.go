package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/config"
	"go.uber.org/zap"
)

type CallbackHandler struct {
	log            *zap.Logger
	Authenticator  *auth.Authenticator
	SessionManager *auth.SessionManager
	Config         *config.Config
}

func NewCallbackHandler(log *zap.Logger, authenticator *auth.Authenticator, sessionManager *auth.SessionManager, config *config.Config) *CallbackHandler {
	return &CallbackHandler{log: log, Authenticator: authenticator, SessionManager: sessionManager, Config: config}
}

func (handler *CallbackHandler) Callback(ctx *fiber.Ctx) error {
	session, _ := handler.SessionManager.SessionStore.Get(ctx)
	state := session.Get("state")
	log.Info(state)

	if ctx.Query("state", "") != state {
		ctx.Response().SetStatusCode(http.StatusBadRequest)
		return ctx.SendString("Bad request")
	}

	code := ctx.Query("code")

	log.Info(code)

	token, err := handler.Authenticator.Exchange(ctx.Context(), code)
	if err != nil {
		return ctx.SendString("Failed to exchange an authorization code for a token.")
	}

	idToken, err := handler.Authenticator.VerifyIDToken(ctx.Context(), token)
	if err != nil {
		return ctx.SendString("Failed to verify ID Token.")
	}

	var profile map[string]interface{}
	if err := idToken.Claims(&profile); err != nil {
		return ctx.SendString(err.Error())
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	if err := session.Save(); err != nil {
		return ctx.SendString(err.Error())
	}

	return ctx.Redirect("/user", http.StatusTemporaryRedirect)
}
