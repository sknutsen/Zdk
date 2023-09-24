package middleware

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func IsAuthenticated(ctx *fiber.Ctx) error {
	if ctx.UserContext() == nil {
		return ctx.Redirect("/", http.StatusSeeOther)
	} else {
		return ctx.Next()
	}
}
