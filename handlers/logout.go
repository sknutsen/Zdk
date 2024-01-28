package handlers

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/config"
)

type LogoutHandler struct {
	Config *config.Config
}

func NewLogoutHandler(config *config.Config) *LogoutHandler {
	return &LogoutHandler{Config: config}
}

func (handler *LogoutHandler) Logout(ctx *fiber.Ctx) error {
	logoutUrl, err := url.Parse(handler.Config.AuthDomain + "/v2/logout")
	if err != nil {
		return err
	}

	log.Println(logoutUrl)

	scheme := "http"
	if ctx.Context().IsTLS() {
		scheme = "https"
	}

	returnTo, err := url.Parse(scheme + "://" + string(ctx.Request().Host()))
	if err != nil {
		return err
	}

	log.Println(returnTo)

	parameters := url.Values{}
	parameters.Add("returnTo", returnTo.String())
	parameters.Add("client_id", handler.Config.AuthClientId)
	logoutUrl.RawQuery = parameters.Encode()

	return ctx.Redirect(logoutUrl.String(), http.StatusTemporaryRedirect)
}
