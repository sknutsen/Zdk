package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sknutsen/Zdk/internal/config"
	"github.com/sknutsen/Zdk/internal/handlers"
	"github.com/sknutsen/Zdk/internal/middleware"
)

func setupTasksRoutes(
	app *fiber.App,
	config *config.Config,
	tasksHandler *handlers.TasksHandler,
	taskCategoriesHandler *handlers.TaskCategoriesHandler,
	scheduledTasksHandler *handlers.ScheduledTasksHandler,
) {
	authMiddleware := middleware.NewJWTMiddleware(config)

	tasksGroup := app.Group("/tasks")
	tasksGroup.Use(authMiddleware.GetAuthMiddleware())
	tasksGroup.Use(middleware.EnsureValidScope)
	tasksGroup.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})
	tasksGroup.Get("/list", tasksHandler.List)
	tasksGroup.Post("/new", tasksHandler.New)
	tasksGroup.Post("/update", tasksHandler.Update)
	tasksGroup.Post("/delete", tasksHandler.Delete)
	tasksGroup.Post("/categories", tasksHandler.Delete)
	// tasksGroup.Post("/addtocategory")
	// tasksGroup.Post("/removefromcategory")

	categoriesGroup := app.Group("/categories")
	categoriesGroup.Use(authMiddleware.GetAuthMiddleware())
	categoriesGroup.Use(middleware.EnsureValidScope)
	categoriesGroup.Get("/list", tasksHandler.List)
	categoriesGroup.Post("/new", tasksHandler.New)
	categoriesGroup.Post("/update", tasksHandler.Update)
	categoriesGroup.Post("/delete", tasksHandler.Delete)
	categoriesGroup.Post("/tasks", tasksHandler.Delete)

	taskCategoriesGroup := app.Group("/taskCategories")
	taskCategoriesGroup.Use(authMiddleware.GetAuthMiddleware())
	taskCategoriesGroup.Use(middleware.EnsureValidScope)
	taskCategoriesGroup.Get("/list", tasksHandler.List)
	taskCategoriesGroup.Post("/new", tasksHandler.New)
	taskCategoriesGroup.Post("/update", tasksHandler.Update)
	taskCategoriesGroup.Post("/delete", tasksHandler.Delete)

	scheduledTasksGroup := app.Group("/scheduledtasks")
	scheduledTasksGroup.Use(authMiddleware.GetAuthMiddleware())
	scheduledTasksGroup.Use(middleware.EnsureValidScope)
	// scheduledTasksGroup.Post("/current")
	scheduledTasksGroup.Get("/list", tasksHandler.List)
	scheduledTasksGroup.Post("/new", tasksHandler.New)
	scheduledTasksGroup.Post("/update", tasksHandler.Update)
	scheduledTasksGroup.Post("/delete", tasksHandler.Delete)
}
