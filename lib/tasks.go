package lib

import (
	"time"

	"github.com/sknutsen/Zdk/models"
)

func CategoryIsSelected(category models.Category, categories []string) bool {
	for _, t := range categories {
		if t == category.CategoryId {
			return true
		}
	}

	return false
}

func TaskIsSelected(task models.Task, tasks []string) bool {
	for _, t := range tasks {
		if t == task.TaskId {
			return true
		}
	}

	return false
}

func BuildCategoryTaskRequest(categoryId string, tasks []string) []models.TaskCategory {
	result := []models.TaskCategory{}

	for _, v := range tasks {
		result = append(result, models.TaskCategory{
			TaskId:      v,
			CategoryId:  categoryId,
			DateCreated: time.Now(),
			DateChanged: time.Now(),
		})
	}

	return result
}

func BuildTaskCategoryRequest(taskId string, categories []string) []models.TaskCategory {
	result := []models.TaskCategory{}

	for _, v := range categories {
		result = append(result, models.TaskCategory{
			TaskId:      taskId,
			CategoryId:  v,
			DateCreated: time.Now(),
			DateChanged: time.Now(),
		})
	}

	return result
}
