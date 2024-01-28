package handlers

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"github.com/sknutsen/Zdk/auth"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/models"
	"github.com/sknutsen/Zdk/view"
	"go.uber.org/zap"
)

type ScheduledTasksHandler struct {
	Ctx            *data.ZdkContext
	SessionManager *auth.SessionManager
	UserHandler    *UserHandler
	Log            *zap.Logger
}

func NewScheduledTasksHandler(
	ctx *data.ZdkContext,
	userHandler *UserHandler,
	sessionManager *auth.SessionManager,
) *ScheduledTasksHandler {
	return &ScheduledTasksHandler{
		Ctx:            ctx,
		UserHandler:    userHandler,
		SessionManager: sessionManager,
	}
}

func (handler *ScheduledTasksHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOScheduledTaskNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	object := models.ScheduledTask{
		ScheduledTaskId: uuid.NewString(),
		TaskId:          request.TaskId,
		IsComplete:      false,
		Date:            time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
		DateCreated:     time.Now(),
		DateChanged:     time.Now(),
	}

	handler.Ctx.DB.Create(object)

	return ctx.SendString("")
}

func (handler *ScheduledTasksHandler) Schedule(c *fiber.Ctx) error {
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

	state := models.GetTasksState(profile)

	var task *models.ScheduledTask
	var taskName string

	year, month, day := time.Now().Date()

	handler.Ctx.DB.Where("tasks.user_id = ? AND scheduled_tasks.date >= ? AND scheduled_tasks.date <= ?", userId, time.Date(year, month, day, 0, 0, 0, 0, time.UTC), time.Date(year, month, day, 23, 59, 59, 0, time.UTC)).Joins("INNER JOIN tasks ON tasks.task_id = scheduled_tasks.task_id").First(&task)

	if task != nil && task.ScheduledTaskId != "" {
		log.Debug(task)

		handler.Ctx.DB.Model(&models.Task{}).Select("name").Where("user_id = ? AND task_id = ?", userId, task.TaskId).First(&taskName)

		state.CurrentScheduledItem = &models.DTOScheduledTaskListResponseData{
			ScheduledTaskId: task.ScheduledTaskId,
			TaskId:          task.TaskId,
			Name:            taskName,
			IsComplete:      task.IsComplete,
			Date:            task.Date,
		}
	}

	var history []models.DTOScheduledTaskListResponseData

	handler.Ctx.DB.Model(&models.ScheduledTask{}).Select("scheduled_tasks.scheduled_task_id, tasks.task_id, tasks.name, scheduled_tasks.is_complete, scheduled_tasks.date").Where("tasks.user_id = ? AND scheduled_tasks.date <= ?", userId, time.Date(year, month, day, 0, 0, 0, 0, time.UTC)).Joins("INNER JOIN tasks ON tasks.task_id = scheduled_tasks.task_id").Find(&history)

	state.TaskHistory = models.GroupScheduledTasks(history)

	return view.Render(c, view.Schedule(state))
}

func (handler *ScheduledTasksHandler) GetTask(c *fiber.Ctx) error {
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

	var tasks []models.Task

	handler.Ctx.DB.Where("user_id = ?", userId).Find(&tasks)

	if len(tasks) > 0 {
		var i = 0

		if len(tasks) > 1 {
			i = rand.Intn(len(tasks) - 1)
		}

		scheduledTask := models.ScheduledTask{
			ScheduledTaskId: uuid.NewString(),
			TaskId:          tasks[i].TaskId,
			IsComplete:      false,
			Date:            time.Now(),
			DateCreated:     time.Now(),
			DateChanged:     time.Now(),
		}

		handler.Ctx.DB.Create(&scheduledTask)

		currentTask := &models.DTOScheduledTaskListResponseData{
			ScheduledTaskId: scheduledTask.ScheduledTaskId,
			TaskId:          scheduledTask.TaskId,
			Name:            tasks[i].Name,
			IsComplete:      false,
			Date:            scheduledTask.Date,
		}

		return view.Render(c, view.CurrentTask(currentTask))
	}

	return view.Render(c, view.CurrentTask(nil))
}

func (handler *ScheduledTasksHandler) Complete(ctx *fiber.Ctx) error {
	var request models.DTOScheduledTaskCompleteRequest

	if err := ctx.BodyParser(&request); err != nil {
		log.Error(err.Error())
		return err
	}

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

	var task models.ScheduledTask

	handler.Ctx.DB.Where("scheduled_task_id = ?", request.ScheduledTaskId).First(&task)

	if task.ScheduledTaskId != "" {
		if !task.IsComplete {
			task.IsComplete = true
			handler.Ctx.DB.Save(&task)
		}

		var taskName string
		handler.Ctx.DB.Model(&models.Task{}).Select("name").Where("task_id = ?", task.TaskId).First(&taskName)

		currentTask := &models.DTOScheduledTaskListResponseData{
			ScheduledTaskId: task.ScheduledTaskId,
			TaskId:          task.TaskId,
			Name:            taskName,
			IsComplete:      true,
			Date:            task.Date,
		}
		return view.Render(ctx, view.CurrentTask(currentTask))
	}

	return view.Render(ctx, view.CurrentTask(nil))
}
