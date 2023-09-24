package routes

import (
	"encoding/gob"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sknutsen/Zdk/internal/auth"
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/handlers"
	"github.com/sknutsen/Zdk/internal/middleware"
	"github.com/sknutsen/Zdk/internal/models"
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
	taskCategoriesHandler *handlers.TaskCategoriesHandler,
	scheduledTasksHandler *handlers.ScheduledTasksHandler,
	sessionManager *auth.SessionManager,
	authenticator *auth.Authenticator,
) {
	gob.Register(map[string]interface{}{})
	gob.Register(models.UserProfile{})
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

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
	app.Static("/", "./wwwroot")

	// app.Use(logger.New(logger.Config{
	// 	Output: log.,
	// }))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/login", loginHandler.Login)
	app.Get("/callback", callbackHandler.Callback)
	app.Get("/user", middleware.IsAuthenticated, func(c *fiber.Ctx) error {
		session, err := sessionManager.SessionStore.Get(c)
		if err != nil {
			log.Error(err.Error())
			return err
		}

		profile := session.Get("profile")

		// var userProfile models.UserProfile
		if profile != nil {
			// for k, v := range profile.(map[string]interface{}) {
			// 	log.Debug("profile", zap.String("Key", k))
			// 	log.Debug("profile", zap.Any("Value", v))
			// }
			// userProfile = profile.(models.UserProfile)
			// log.Debug("userProfile", zap.String("user id", userProfile.UserId))
		} else {
			log.Debug("profile is null")
		}
		// log.Info(profile.Nickname)

		return c.Render("user", profile, "layouts/main")
		// return c.Render("user", fiber.Map{
		// 	"nickname": profile,
		// }, "layouts/main")
	})
	app.Get("/logout", logoutHandler.Logout)

	setupShoppingListRoutes(app, config, shoppingListItemsHandler, shoppingListsHandler)
	setupTasksRoutes(app, config, tasksHandler, taskCategoriesHandler, scheduledTasksHandler)

	log.Fatal(app.Listen(address).Error())
}
