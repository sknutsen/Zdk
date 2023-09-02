package routes

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/handlers"
	"github.com/sknutsen/Zdk/internal/middleware"
)

type IRouteHandler interface {
	BindRoutes()
	List()
	New()
	Update()
	Delete()
}

func NewRouter(config *config.Config, shoppingListItemsHandler *handlers.ShoppingListItemsHandler, shoppingListsHandler *handlers.ShoppingListItemsHandler) {
	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})

	port := os.Getenv("PORT")
	var hostName string

	if port == "" {
		port = "8080"
		hostName = ""
	} else {
		hostName = "0.0.0.0"
	}

	address := fmt.Sprintf("%s:%s", hostName, port)

	app.Static("/", "./wwwroot")

	// app.Use()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	authMiddleware := middleware.NewJWTMiddleware(config)

	shoppingListGroup := app.Group("/shoppinglists")
	shoppingListGroup.Use(authMiddleware.GetAuthMiddleware())
	shoppingListGroup.Use(middleware.EnsureValidScope)
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

	log.Fatal(app.Listen(address))
}
