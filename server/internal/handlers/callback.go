package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sknutsen/Zdk/internal/auth"
	"github.com/sknutsen/Zdk/internal/config"
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

type profile struct {
	iss        string
	sub        string
	name       string
	aud        string
	sid        string
	nickname   string
	picture    string
	iat        string
	exp        string
	updated_at string
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
	// var profile profile
	if err := idToken.Claims(&profile); err != nil {
		return ctx.SendString(err.Error())
	}

	session.Set("access_token", token.AccessToken)
	session.Set("profile", profile)
	// session.Set("profile", models.UserProfile{
	// 	UserId:    profile.sub,
	// 	Name:      profile.name,
	// 	Nickname:  profile.nickname,
	// 	Picture:   profile.picture,
	// 	Iat:       profile.iat,
	// 	Exp:       profile.exp,
	// 	UpdatedAt: profile.updated_at,
	// })
	// handler.SessionManager.SessionStore.Storage.Set("profile", profile, time.Hour*24)
	if err := session.Save(); err != nil {
		return ctx.SendString(err.Error())
	}

	// Redirect to logged in page.
	return ctx.Redirect("/user", http.StatusTemporaryRedirect)
}
