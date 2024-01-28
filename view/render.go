package view

import (
	"context"

	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
)

func Render(c *fiber.Ctx, cmp templ.Component) error {
	c.Set(fiber.HeaderContentType, fiber.MIMETextHTML)
	return cmp.Render(context.Background(), c.Response().BodyWriter())
}
