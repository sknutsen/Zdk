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

func setupWorkoutsRoutes(
	app *fiber.App,
	config *config.Config,
	workoutHandler *handlers.WorkoutHandler,
	sessionManager *auth.SessionManager,
	mw *middleware.Middleware,
	log *zap.Logger,
) {
	workoutsGroup := app.Group("/workouts")
	// tasksGroup.Use(mw.GetAuthMiddleware())
	// tasksGroup.Use(middleware.EnsureValidScope)
	workoutsGroup.Get("/", mw.IsAuthenticated, func(c *fiber.Ctx) error {
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

		userId, err := workoutHandler.UserHandler.GetUserId(c)
		if err != nil {
			return err
		}

		var tasks []models.Workout

		workoutHandler.Ctx.DB.Where("user_id = ?", userId).Find(&tasks)

		state := models.GetWorkoutState(profile)

		state.Workouts = tasks

		return view.Render(c, view.WorkoutsPage(state))
	})
	workoutsGroup.Get("/list", mw.IsAuthenticated, workoutHandler.List)
	workoutsGroup.Get("/create", mw.IsAuthenticated, workoutHandler.Create)
	workoutsGroup.Get("/edit/:id", mw.IsAuthenticated, workoutHandler.Edit)
	workoutsGroup.Post("/new", mw.IsAuthenticated, workoutHandler.New)
	workoutsGroup.Post("/update", mw.IsAuthenticated, workoutHandler.Update)
	workoutsGroup.Get("/delete/:id", mw.IsAuthenticated, workoutHandler.Delete)

	workoutsGroup.Get("/edit/:wid/exercises/list", mw.IsAuthenticated, workoutHandler.ListExercises)
	workoutsGroup.Get("/edit/:wid/exercises/create", mw.IsAuthenticated, workoutHandler.CreateExercise)
	workoutsGroup.Get("/edit/:wid/exercises/edit/:id", mw.IsAuthenticated, workoutHandler.EditExercise)
	workoutsGroup.Post("/edit/:wid/exercises/new", mw.IsAuthenticated, workoutHandler.NewExercise)
	workoutsGroup.Post("/edit/:wid/exercises/update", mw.IsAuthenticated, workoutHandler.UpdateExercise)
	workoutsGroup.Get("/edit/:wid/exercises/delete/:id", mw.IsAuthenticated, workoutHandler.DeleteExercise)
}
