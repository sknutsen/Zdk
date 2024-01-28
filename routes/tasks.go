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

func setupTasksRoutes(
	app *fiber.App,
	config *config.Config,
	tasksHandler *handlers.TasksHandler,
	taskCategoriesHandler *handlers.TaskCategoriesHandler,
	categoriesHandler *handlers.CategoriesHandler,
	scheduledTasksHandler *handlers.ScheduledTasksHandler,
	sessionManager *auth.SessionManager,
	mw *middleware.Middleware,
	log *zap.Logger,
) {
	tasksGroup := app.Group("/tasks")
	// tasksGroup.Use(mw.GetAuthMiddleware())
	// tasksGroup.Use(middleware.EnsureValidScope)
	tasksGroup.Get("/", mw.IsAuthenticated, func(c *fiber.Ctx) error {
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

		userId, err := tasksHandler.UserHandler.GetUserId(c)
		if err != nil {
			return err
		}

		var tasks []models.Task

		tasksHandler.Ctx.DB.Where("user_id = ?", userId).Find(&tasks)

		state := models.GetTasksState(profile)

		state.Tasks = tasks

		return view.Render(c, view.TasksPage(state))
	})
	tasksGroup.Get("/create", mw.IsAuthenticated, tasksHandler.Create)
	tasksGroup.Get("/edit/:id", mw.IsAuthenticated, tasksHandler.Edit)
	tasksGroup.Post("/new", mw.IsAuthenticated, tasksHandler.New)
	tasksGroup.Post("/update", mw.IsAuthenticated, tasksHandler.Update)
	tasksGroup.Get("/delete/:id", mw.IsAuthenticated, tasksHandler.Delete)

	categoriesGroup := app.Group("/categories")
	// categoriesGroup.Use(mw.GetAuthMiddleware())
	// categoriesGroup.Use(middleware.EnsureValidScope)
	categoriesGroup.Get("/", mw.IsAuthenticated, func(c *fiber.Ctx) error {
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

		userId, err := categoriesHandler.UserHandler.GetUserId(c)
		if err != nil {
			return err
		}

		var categories []models.Category

		categoriesHandler.Ctx.DB.Where("user_id = ?", userId).Find(&categories)

		state := models.GetTasksState(profile)

		state.Categories = categories

		return view.Render(c, view.CategoriesPage(state))
	})
	categoriesGroup.Get("/create", mw.IsAuthenticated, categoriesHandler.Create)
	categoriesGroup.Get("/edit/:id", mw.IsAuthenticated, categoriesHandler.Edit)
	categoriesGroup.Post("/new", mw.IsAuthenticated, categoriesHandler.New)
	categoriesGroup.Post("/update", mw.IsAuthenticated, categoriesHandler.Update)
	categoriesGroup.Get("/delete/:id", mw.IsAuthenticated, categoriesHandler.Delete)

	scheduledTasksGroup := app.Group("/scheduled-tasks")
	// scheduledTasksGroup.Use(mw.GetAuthMiddleware())
	// scheduledTasksGroup.Use(middleware.EnsureValidScope)
	scheduledTasksGroup.Get("/", mw.IsAuthenticated, scheduledTasksHandler.Schedule)
	scheduledTasksGroup.Get("/get-task", mw.IsAuthenticated, scheduledTasksHandler.GetTask)
	scheduledTasksGroup.Post("/complete", mw.IsAuthenticated, scheduledTasksHandler.Complete)
}
