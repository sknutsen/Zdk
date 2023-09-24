package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/handlers"
	"github.com/sknutsen/Zdk/internal/middleware"
)

func setupShoppingListRoutes(
	app *fiber.App,
	config *config.Config,
	shoppingListItemsHandler *handlers.ShoppingListItemsHandler,
	shoppingListsHandler *handlers.ShoppingListsHandler,
) {
	authMiddleware := middleware.NewJWTMiddleware(config)

	shoppingListGroup := app.Group("/shoppinglists")
	shoppingListGroup.Use(authMiddleware.GetAuthMiddleware())
	// shoppingListGroup.Use(middleware.EnsureValidScope)
	shoppingListGroup.Use(middleware.HTTPMiddleware)
	shoppingListGroup.Get(("/"), func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})
	shoppingListGroup.Get(("/list"), shoppingListsHandler.List)
	shoppingListGroup.Post("/new", shoppingListsHandler.New)
	shoppingListGroup.Post("/update", shoppingListsHandler.Update)
	shoppingListGroup.Post("/delete", shoppingListsHandler.Delete)

	shoppingListItemsGroup := app.Group("/shoppinglistitems")
	shoppingListItemsGroup.Use(authMiddleware.GetAuthMiddleware())
	shoppingListItemsGroup.Use(middleware.EnsureValidScope)
	shoppingListItemsGroup.Get(("/list"), shoppingListItemsHandler.List)
	shoppingListItemsGroup.Post("/new", shoppingListItemsHandler.New)
	shoppingListItemsGroup.Post("/update", shoppingListItemsHandler.Update)
	shoppingListItemsGroup.Post("/delete", shoppingListItemsHandler.Delete)
}
