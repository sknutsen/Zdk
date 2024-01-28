package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/config"
	"github.com/sknutsen/Zdk/handlers"
	"github.com/sknutsen/Zdk/middleware"
	"github.com/sknutsen/Zdk/models"
	"github.com/sknutsen/Zdk/view"
	"go.uber.org/zap"
)

func setupShoppingListRoutes(
	app *fiber.App,
	config *config.Config,
	shoppingListsHandler *handlers.ShoppingListsHandler,
	shoppingListItemsHandler *handlers.ShoppingListItemsHandler,
	sessionManager *auth.SessionManager,
	mw *middleware.Middleware,
	log *zap.Logger,
) {
	shoppingListGroup := app.Group("/shoppinglists")
	// shoppingListGroup.Use(mw.GetAuthMiddleware())
	// shoppingListGroup.Use(middleware.EnsureValidScope)
	// shoppingListGroup.Use(middleware.HTTPMiddleware)
	shoppingListGroup.Get(("/"), mw.IsAuthenticated, func(c *fiber.Ctx) error {
		session, err := sessionManager.SessionStore.Get(c)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		profile := session.Get("profile")

		if profile == nil {
			log.Debug("profile is null")
			return c.Redirect("/")
		}

		state := models.GetShoppingListsState(profile)

		return view.Render(c, view.ShoppingListsPage(state))
	})
	shoppingListGroup.Get(("/list"), shoppingListsHandler.List)
	shoppingListGroup.Post("/new", shoppingListsHandler.New)
	shoppingListGroup.Post("/update", shoppingListsHandler.Update)
	shoppingListGroup.Post("/delete", shoppingListsHandler.Delete)

	shoppingListItemsGroup := app.Group("/shoppinglistitems")
	// shoppingListItemsGroup.Use(mw.GetAuthMiddleware())
	// shoppingListItemsGroup.Use(middleware.EnsureValidScope)
	shoppingListItemsGroup.Get(("/list"), shoppingListItemsHandler.List)
	shoppingListItemsGroup.Post("/new", shoppingListItemsHandler.New)
	shoppingListItemsGroup.Post("/update", shoppingListItemsHandler.Update)
	shoppingListItemsGroup.Post("/delete", shoppingListItemsHandler.Delete)
}
