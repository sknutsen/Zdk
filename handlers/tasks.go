package handlers

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/lib"
	"github.com/sknutsen/Zdk/models"
	"github.com/sknutsen/Zdk/view"
)

type TasksHandler struct {
	Ctx            *data.ZdkContext
	SessionManager *auth.SessionManager
	UserHandler    *UserHandler
}

func NewTasksHandler(
	ctx *data.ZdkContext,
	userHandler *UserHandler,
	sessionManager *auth.SessionManager,
) *TasksHandler {
	return &TasksHandler{
		Ctx:            ctx,
		SessionManager: sessionManager,
		UserHandler:    userHandler,
	}
}

func (handler *TasksHandler) List(ctx *fiber.Ctx) error {
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

	var tasks []models.Task

	handler.Ctx.DB.Where("user_id = ?", userId).Find(&tasks)

	return view.Render(ctx, view.ListTasks(tasks))
}

func (handler *TasksHandler) New(ctx *fiber.Ctx) error {
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

	var request models.DTOTaskNewRequest

	if err := ctx.BodyParser(&request); err != nil {
		log.Error(err.Error())
		return err
	}

	userId, err := handler.UserHandler.GetUserId(ctx)
	if err != nil {
		return err
	}

	object := models.Task{
		TaskId:      uuid.NewString(),
		UserId:      userId,
		Name:        request.Name,
		DateCreated: time.Now(),
		DateChanged: time.Now(),
	}

	handler.Ctx.DB.Create(object)

	var taskCategory = lib.BuildTaskCategoryRequest(object.TaskId, request.Categories)

	if len(taskCategory) > 0 {
		handler.Ctx.DB.Create(&taskCategory)
	}

	return ctx.Redirect("/tasks")
}

func (handler *TasksHandler) Update(ctx *fiber.Ctx) error {
	var request models.DTOTaskUpdateRequest

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

	var task models.Task

	result := handler.Ctx.DB.Where("user_id = ? AND task_id = ?", userId, request.TaskId).First(&task)

	if result.Error != nil {
		log.Error(result.Error)

		return ctx.Redirect(fmt.Sprintf("/tasks/edit/%s", task.TaskId))
	}

	object := models.Task{
		TaskId:      request.TaskId,
		UserId:      userId,
		Name:        request.Name,
		DateChanged: time.Now(),
	}

	handler.Ctx.DB.Save(object)

	var taskCategories = lib.BuildTaskCategoryRequest(request.TaskId, request.Categories)

	listUpdateTaskCategories(*handler.Ctx.DB, userId, taskCategories)

	return ctx.Redirect("/tasks/list")
}

func (handler *TasksHandler) Delete(ctx *fiber.Ctx) error {
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

	object := models.Task{
		TaskId: id,
	}

	handler.Ctx.DB.Where("user_id = ?", userId).Delete(object)
	handler.Ctx.DB.Where("task_id = ?", id).Delete(&models.TaskCategory{})

	return ctx.Redirect("/tasks/list")
}

func (handler *TasksHandler) Create(c *fiber.Ctx) error {
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

	var categories []models.Category
	handler.Ctx.DB.Where("user_id = ?", userId).Find(&categories)

	return view.Render(c, view.CreateTask(categories))
}

func (handler *TasksHandler) Edit(c *fiber.Ctx) error {
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

	var task *models.Task

	handler.Ctx.DB.Where("task_id = ?", id).First(&task)

	if task == nil {
		log.Debug("task does not exist")
		c.Redirect("/tasks")
	}

	var selectedCategories []string
	var categories []models.Category

	userId, err := handler.UserHandler.GetUserId(c)
	if err != nil {
		return err
	}

	handler.Ctx.DB.Model(&models.TaskCategory{}).Select("category_id").Joins("INNER JOIN tasks ON tasks.task_id = task_categories.task_id").Where("tasks.user_id = ?", userId).Find(&selectedCategories)
	handler.Ctx.DB.Where("user_id = ?", userId).Find(&categories)

	return view.Render(c, view.EditTask(
		&models.DTOTaskUpdateRequest{
			TaskId:     task.TaskId,
			Name:       task.Name,
			Categories: selectedCategories,
		},
		categories,
	))
}
