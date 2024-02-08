package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/models"
	"github.com/sknutsen/Zdk/view"
)

type WorkoutHandler struct {
	Ctx            *data.ZdkContext
	SessionManager *auth.SessionManager
	UserHandler    *UserHandler
}

func NewWorkoutHandler(
	ctx *data.ZdkContext,
	userHandler *UserHandler,
	sessionManager *auth.SessionManager,
) *WorkoutHandler {
	return &WorkoutHandler{
		Ctx:            ctx,
		SessionManager: sessionManager,
		UserHandler:    userHandler,
	}
}

func (handler *WorkoutHandler) List(ctx *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return ctx.Redirect("/")
	}

	userId, err := handler.UserHandler.GetUserId(ctx)
	if err != nil {
		return err
	}

	var tasks []models.Workout

	handler.Ctx.DB.Where("user_id = ?", userId).Find(&tasks)

	return view.Render(ctx, view.ListWorkouts(tasks))
}

func (handler *WorkoutHandler) New(ctx *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return ctx.Redirect("/")
	}

	var request models.DTOWorkoutNewRequest

	if err := ctx.BodyParser(&request); err != nil {
		log.Error(err.Error())
		return err
	}

	userId, err := handler.UserHandler.GetUserId(ctx)
	if err != nil {
		return err
	}

	object := models.Workout{
		WorkoutId:   uuid.NewString(),
		UserId:      userId,
		Name:        request.Name,
		Date:        request.Date,
		Description: request.Description,
		DateCreated: time.Now(),
		DateChanged: time.Now(),
	}

	handler.Ctx.DB.Create(object)

	return ctx.Redirect("/workouts/list")
}

func (handler *WorkoutHandler) Update(ctx *fiber.Ctx) error {
	var request models.DTOWorkoutUpdateRequest

	session, err := handler.SessionManager.SessionStore.Get(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return ctx.Redirect("/")
	}

	if err := ctx.BodyParser(&request); err != nil {
		log.Error(err.Error())
		return err
	}

	userId, err := handler.UserHandler.GetUserId(ctx)
	if err != nil {
		return err
	}

	var task models.Workout

	result := handler.Ctx.DB.Where("user_id = ? AND workout_id = ?", userId, request.WorkoutId).First(&task)

	if result.Error != nil {
		log.Error(result.Error)

		return ctx.Redirect(fmt.Sprintf("/workouts/edit/%s", task.WorkoutId))
	}

	object := models.Workout{
		WorkoutId:   request.WorkoutId,
		UserId:      userId,
		Name:        request.Name,
		Date:        request.Date,
		Description: request.Description,
		DateChanged: time.Now(),
	}

	handler.Ctx.DB.Save(object)

	return ctx.Redirect("/workouts/list")
}

func (handler *WorkoutHandler) Delete(ctx *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(ctx)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return ctx.Redirect("/")
	}

	userId, err := handler.UserHandler.GetUserId(ctx)
	if err != nil {
		return err
	}

	id := ctx.Params("id")

	object := models.Workout{
		WorkoutId: id,
	}

	handler.Ctx.DB.Where("user_id = ?", userId).Delete(object)

	return ctx.Redirect("/workouts/list")
}

func (handler *WorkoutHandler) Create(c *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(c)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return c.Redirect("/")
	}

	return view.Render(c, view.CreateWorkout())
}

func (handler *WorkoutHandler) Edit(c *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(c)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return c.Redirect("/")
	}

	id := c.Params("id")

	var workout *models.Workout

	handler.Ctx.DB.Where("workout_id = ?", id).First(&workout)

	if workout == nil {
		log.Debug("workout does not exist")
		c.Redirect("/workouts")
	}

	var exercises []models.Exercise

	userId, err := handler.UserHandler.GetUserId(c)
	if err != nil {
		return err
	}

	handler.Ctx.DB.Where("user_id = ? AND workout_id = ?", userId, id).Find(&exercises)

	return view.Render(c, view.EditWorkout(
		&models.DTOWorkoutUpdateRequest{
			WorkoutId:   workout.WorkoutId,
			Name:        workout.Name,
			Date:        workout.Date,
			Description: workout.Description,
			Exercises:   exercises,
		},
	))
}

func (handler *WorkoutHandler) ListExercises(c *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(c)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return c.Redirect("/")
	}

	userId, err := handler.UserHandler.GetUserId(c)
	if err != nil {
		return err
	}

	wid := c.Params("wid")

	var tasks []models.Exercise

	handler.Ctx.DB.Where("user_id = ? AND workout_id = ?", userId, wid).Find(&tasks)

	return view.Render(c, view.ListExercises(wid, tasks))
}

func (handler *WorkoutHandler) NewExercise(c *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(c)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return c.Redirect("/")
	}

	var request models.DTOExerciseNewRequest

	if err := c.BodyParser(&request); err != nil {
		log.Error(err.Error())
		return err
	}

	userId, err := handler.UserHandler.GetUserId(c)
	if err != nil {
		return err
	}

	wid := c.Params("wid")

	object := models.Exercise{
		ExerciseId:    uuid.NewString(),
		UserId:        userId,
		WorkoutId:     wid,
		ExerciseDefId: request.ExerciseDefId,
		EquipmentId:   request.EquipmentId,
		Hours:         request.Hours,
		Minutes:       request.Minutes,
		Seconds:       request.Seconds,
		Units:         request.Units,
		UnitTypeId:    request.UnitTypeId,
		Sets:          request.Sets,
		Weight:        request.Weight,
		Name:          request.Name,
		DateCreated:   time.Now(),
		DateChanged:   time.Now(),
	}

	handler.Ctx.DB.Create(object)

	return c.Redirect(fmt.Sprintf("/workouts/edit/%s/exercises/list", wid))
}

func (handler *WorkoutHandler) UpdateExercise(c *fiber.Ctx) error {
	var request models.DTOExerciseUpdateRequest

	session, err := handler.SessionManager.SessionStore.Get(c)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return c.Redirect("/")
	}

	if err := c.BodyParser(&request); err != nil {
		log.Error(err.Error())
		return err
	}

	userId, err := handler.UserHandler.GetUserId(c)
	if err != nil {
		return err
	}

	wid := c.Params("wid")

	var exercise models.Exercise

	result := handler.Ctx.DB.Where("user_id = ? AND exercise_id = ?", userId, request.ExerciseId).First(&exercise)

	if result.Error != nil {
		log.Error(result.Error)

		return c.Redirect(fmt.Sprintf("/workouts/edit/%s/exercises/list", wid))
	}

	object := models.Exercise{
		ExerciseId:    request.ExerciseId,
		WorkoutId:     wid,
		ExerciseDefId: request.ExerciseDefId,
		EquipmentId:   request.EquipmentId,
		Hours:         request.Hours,
		Minutes:       request.Minutes,
		Seconds:       request.Seconds,
		Units:         request.Units,
		UnitTypeId:    request.UnitTypeId,
		Sets:          request.Sets,
		Weight:        request.Weight,
		Name:          request.Name,
		UserId:        userId,
		DateChanged:   time.Now(),
	}

	handler.Ctx.DB.Save(object)

	return c.Redirect(fmt.Sprintf("/workouts/edit/%s/exercises/list", wid))
}

func (handler *WorkoutHandler) DeleteExercise(c *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(c)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return c.Redirect("/")
	}

	userId, err := handler.UserHandler.GetUserId(c)
	if err != nil {
		return err
	}

	wid := c.Params("wid")
	id := c.Params("id")

	object := models.Exercise{
		ExerciseId: id,
	}

	handler.Ctx.DB.Where("user_id = ?", userId).Delete(object)

	return c.Redirect(fmt.Sprintf("/workouts/edit/%s/exercises/list", wid))
}

func (handler *WorkoutHandler) CreateExercise(c *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(c)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return c.Redirect("/")
	}

	wid := c.Params("wid")

	var exerciseDefs []models.ExerciseDef
	var equipment []models.Equipment
	var unitTypes []models.UnitType

	handler.Ctx.DB.Find(&exerciseDefs)
	handler.Ctx.DB.Find(&equipment)
	handler.Ctx.DB.Find(&unitTypes)

	return view.Render(c, view.CreateExercise(wid, exerciseDefs, equipment, unitTypes))
}

func (handler *WorkoutHandler) EditExercise(c *fiber.Ctx) error {
	session, err := handler.SessionManager.SessionStore.Get(c)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	profile := session.Get("profile")

	if profile == nil {
		log.Debug("profile is null")
		return c.Redirect("/")
	}

	wid := c.Params("wid")
	id := c.Params("id")

	var exercise *models.Exercise

	handler.Ctx.DB.Where("exercise_id = ?", id).First(&exercise)

	if exercise == nil {
		log.Debug("workout does not exist")
		c.Redirect(fmt.Sprintf("/workouts/edit/%s", wid))
	}

	var exerciseDefs []models.ExerciseDef
	var equipment []models.Equipment
	var unitTypes []models.UnitType

	handler.Ctx.DB.Find(&exerciseDefs)
	handler.Ctx.DB.Find(&equipment)
	handler.Ctx.DB.Find(&unitTypes)

	return view.Render(c, view.EditExercise(
		&models.DTOExerciseUpdateRequest{
			ExerciseId:    exercise.ExerciseId,
			ExerciseDefId: exercise.ExerciseDefId,
			EquipmentId:   exercise.EquipmentId,
			WorkoutId:     exercise.WorkoutId,
			Name:          exercise.Name,
			Hours:         exercise.Hours,
			Minutes:       exercise.Minutes,
			Seconds:       exercise.Seconds,
			Units:         exercise.Units,
			UnitTypeId:    exercise.UnitTypeId,
			Sets:          exercise.Sets,
			Weight:        exercise.Weight,
		},
		exerciseDefs,
		equipment,
		unitTypes,
	))
}
