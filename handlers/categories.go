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

type CategoriesHandler struct {
	Ctx            *data.ZdkContext
	SessionManager *auth.SessionManager
	UserHandler    *UserHandler
}

func NewCategoriesHandler(
	ctx *data.ZdkContext,
	userHandler *UserHandler,
	sessionManager *auth.SessionManager,
) *CategoriesHandler {
	return &CategoriesHandler{
		Ctx:            ctx,
		SessionManager: sessionManager,
		UserHandler:    userHandler,
	}
}

func (handler *CategoriesHandler) List(ctx *fiber.Ctx) error {
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

	var categories []models.Category

	handler.Ctx.DB.Where("user_id = ?", userId).Find(&categories)

	return view.Render(ctx, view.ListCategories(categories))
}

func (handler *CategoriesHandler) New(ctx *fiber.Ctx) error {
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

	var request models.DTOCategoryNewRequest

	if err := ctx.BodyParser(&request); err != nil {
		log.Error(err.Error())
		return err
	}

	userId, err := handler.UserHandler.GetUserId(ctx)
	if err != nil {
		return err
	}

	object := models.Category{
		CategoryId:  uuid.NewString(),
		UserId:      userId,
		Name:        request.Name,
		DateCreated: time.Now(),
		DateChanged: time.Now(),
	}

	handler.Ctx.DB.Create(object)

	var taskCategory = lib.BuildCategoryTaskRequest(object.CategoryId, request.Tasks)

	handler.Ctx.DB.Create(&taskCategory)

	return ctx.Redirect("/categories/list")
}

func (handler *CategoriesHandler) Update(ctx *fiber.Ctx) error {
	var request models.DTOCategoryUpdateRequest

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

	var category models.Category

	result := handler.Ctx.DB.Where("user_id = ? AND category_id = ?", userId, request.CategoryId).First(&category)

	if result.Error != nil {
		log.Error(result.Error)

		return ctx.Redirect(fmt.Sprintf("/categories/edit/%s", category.CategoryId))
	}

	object := models.Category{
		CategoryId:  request.CategoryId,
		UserId:      userId,
		Name:        request.Name,
		DateChanged: time.Now(),
	}

	handler.Ctx.DB.Save(object)

	var taskCategories = lib.BuildCategoryTaskRequest(request.CategoryId, request.Tasks)

	listUpdateTaskCategories(*handler.Ctx.DB, userId, taskCategories)

	return ctx.Redirect("/categories/list")
}

func (handler *CategoriesHandler) Delete(ctx *fiber.Ctx) error {
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

	object := models.Category{
		CategoryId: id,
	}

	handler.Ctx.DB.Where("user_id = ?", userId).Delete(object)
	handler.Ctx.DB.Where("user_id = ? AND category_id = ?", userId, id).Delete(&models.TaskCategory{})

	return ctx.Redirect("/categories/list")
}

func (handler *CategoriesHandler) Create(c *fiber.Ctx) error {
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

	return view.Render(c, view.CreateCategory(tasks))
}

func (handler *CategoriesHandler) Edit(c *fiber.Ctx) error {
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

	var category *models.Category

	handler.Ctx.DB.Where("category_id = ?", id).First(&category)

	if category == nil {
		log.Debug("task does not exist")
		c.Redirect("/categories")
	}

	var selectedTasks []string
	var tasks []models.Task

	userId, err := handler.UserHandler.GetUserId(c)
	if err != nil {
		return err
	}

	handler.Ctx.DB.Model(&models.TaskCategory{}).Select("task_id").Joins("INNER JOIN categories ON categories.category_id = task_categories.category_id").Where("categories.user_id = ?", userId).Find(&selectedTasks)
	handler.Ctx.DB.Where("user_id = ?", userId).Find(&tasks)

	return view.Render(c, view.EditCategory(
		&models.DTOCategoryUpdateRequest{
			CategoryId: category.CategoryId,
			Name:       category.Name,
			Tasks:      selectedTasks,
		},
		tasks,
	))
}
