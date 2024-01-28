package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func (mw *Middleware) IsAuthenticated(ctx *fiber.Ctx) error {
	session, err := mw.SessionManager.SessionStore.Get(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return ctx.Redirect("/login")
	}

	return ctx.Next()
}
