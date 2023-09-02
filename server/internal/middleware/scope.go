package middleware

import (
	"net/http"

	jwtmiddleware "github.com/auth0/go-jwt-middleware/v2"
	"github.com/auth0/go-jwt-middleware/v2/validator"
	"github.com/gofiber/fiber/v2"
)

func EnsureValidScope(ctx *fiber.Ctx) error {
	ctx.Response().Header.Set("Access-Control-Allow-Credentials", "true")
	ctx.Response().Header.Set("Access-Control-Allow-Origin", "http://localhost:3000")
	ctx.Response().Header.Set("Access-Control-Allow-Headers", "Authorization")

	token := ctx.Context().Value(jwtmiddleware.ContextKey{}).(*validator.ValidatedClaims)

	claims := token.CustomClaims.(*CustomClaims)
	if !claims.HasScope("read:data") {
		ctx.Response().Header.SetStatusCode(http.StatusForbidden)
		return ctx.SendString("Forbidden")
	}

	return ctx.Next()
}
