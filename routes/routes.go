package routes

import (
	"encoding/gob"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/config"
	"github.com/sknutsen/Zdk/handlers"
	"github.com/sknutsen/Zdk/middleware"
	"github.com/sknutsen/Zdk/models"
	"github.com/sknutsen/Zdk/view"
	"go.uber.org/zap"
)

type IRouteHandler interface {
	BindRoutes()
	List()
	New()
	Update()
	Delete()
}

func NewRouter(
	log *zap.Logger,
	config *config.Config,
	callbackHandler *handlers.CallbackHandler,
	loginHandler *handlers.LoginHandler,
	logoutHandler *handlers.LogoutHandler,
	shoppingListItemsHandler *handlers.ShoppingListItemsHandler,
	shoppingListsHandler *handlers.ShoppingListsHandler,
	tasksHandler *handlers.TasksHandler,
	categoriesHandler *handlers.CategoriesHandler,
	taskCategoriesHandler *handlers.TaskCategoriesHandler,
	scheduledTasksHandler *handlers.ScheduledTasksHandler,
	sessionManager *auth.SessionManager,
	authenticator *auth.Authenticator,
	mw *middleware.Middleware,
) {
	gob.Register(map[string]interface{}{})
	gob.Register(models.UserProfile{})

	app := fiber.New(fiber.Config{})

	port := config.Port
	var hostName string

	if port == "" {
		port = "8080"
		hostName = ""
	} else {
		hostName = "0.0.0.0"
	}

	address := fmt.Sprintf("%s:%s", hostName, port)
	log.Info(address)
	app.Static("/", "./assets")
	app.Get("/", func(c *fiber.Ctx) error {
		session, err := sessionManager.SessionStore.Get(c)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		profile := session.Get("profile")

		state := models.GetIndexState(profile)

		return view.Render(c, view.Index(state))
	})

	app.Get("/login", loginHandler.Login)
	app.Get("/callback", callbackHandler.Callback)
	app.Get("/user", mw.IsAuthenticated, func(c *fiber.Ctx) error {
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

		state := models.GetUserState(profile)

		return view.Render(c, view.User(state))
	})
	app.Get("/logout", logoutHandler.Logout)

	setupShoppingListRoutes(app, config, shoppingListsHandler, shoppingListItemsHandler, sessionManager, mw, log)
	setupTasksRoutes(app, config, tasksHandler, taskCategoriesHandler, categoriesHandler, scheduledTasksHandler, sessionManager, mw, log)

	log.Fatal(app.Listen(address).Error())
}
