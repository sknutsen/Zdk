package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/sknutsen/Zdk/data"
	"github.com/sknutsen/Zdk/models"
	"gorm.io/gorm"
)

type TaskCategoriesHandler struct {
	Ctx         *data.ZdkContext
	UserHandler *UserHandler
}

func NewTaskCategoriesHandler(ctx *data.ZdkContext, userHandler *UserHandler) *TaskCategoriesHandler {
	return &TaskCategoriesHandler{Ctx: ctx, UserHandler: userHandler}
}

func (handler *TaskCategoriesHandler) List(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskCategoryListRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *TaskCategoriesHandler) New(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskCategoryNewRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	object := models.TaskCategory{
		TaskCategoryId: uuid.NewString(),
		TaskId:         request.TaskId,
		CategoryId:     request.CategoryId,
		DateCreated:    time.Now(),
		DateChanged:    time.Now(),
	}

	handler.Ctx.DB.Create(object)

	return ctx.SendString("")
}

func (handler *TaskCategoriesHandler) Update(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskCategoryUpdateRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func (handler *TaskCategoriesHandler) Delete(ctx *fiber.Ctx) error {
	request := new(models.DTOTaskCategoryDeleteRequest)

	if err := ctx.BodyParser(request); err != nil {
		return err
	}

	return ctx.SendString("")
}

func buildTaskCategoriesAddList(selected []models.TaskCategory, list []models.TaskCategory) []models.TaskCategory {
	result := []models.TaskCategory{}

	var exists bool
	if len(selected) > 0 {
		for _, c := range selected {
			exists = false

			for _, tc := range list {
				if tc.CategoryId == c.CategoryId && tc.TaskId == c.TaskId {
					exists = true
					break
				}
			}

			if !exists {
				result = append(result, models.TaskCategory{
					TaskCategoryId: uuid.NewString(),
					TaskId:         c.TaskId,
					CategoryId:     c.CategoryId,
					DateCreated:    time.Now(),
					DateChanged:    time.Now(),
				})
			}
		}
	}

	return result
}

func buildTaskCategoriesDeleteList(selected []models.TaskCategory, list []models.TaskCategory) []models.TaskCategory {
	result := []models.TaskCategory{}

	var exists bool

	for _, tc := range list {
		exists = false

		if len(selected) > 0 {
			for _, c := range selected {
				if c.CategoryId == tc.CategoryId && c.TaskId == tc.TaskId {
					exists = true
					break
				}
			}
		}

		if !exists {
			result = append(result, tc)
		}
	}

	return result
}

func listUpdateTaskCategories(db gorm.DB, userId string, selected []models.TaskCategory) {
	var existing []models.TaskCategory

	db.Where("categories.user_id = ?", userId).Joins("INNER JOIN categories ON categories.category_id = task_categories.category_id").Find(&existing)

	addList := buildTaskCategoriesAddList(selected, existing)
	deleteList := buildTaskCategoriesDeleteList(selected, existing)

	if len(addList) > 0 {
		db.Create(&addList)
	}

	if len(deleteList) > 0 {
		db.Delete(&deleteList)
	}
}
