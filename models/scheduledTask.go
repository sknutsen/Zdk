package models

import (
	"time"
)

type ScheduledTask struct {
	ScheduledTaskId string `gorm:"primaryKey"`
	TaskId          string
	IsComplete      bool
	Date            time.Time
	DateCreated     time.Time
	DateChanged     time.Time
}

type ScheduledTasksByYear struct {
	Year         int
	TasksByMonth []ScheduledTasksByMonth
}

type ScheduledTasksByMonth struct {
	Month       time.Month
	TaskHistory []DTOScheduledTaskListResponseData
}

func GroupScheduledTasks(st []DTOScheduledTaskListResponseData) []ScheduledTasksByYear {
	hist := []ScheduledTasksByYear{}
	var yearExists bool
	var monthExists bool

	for _, t := range st {
		yearExists = false
		monthExists = false

		for _, y := range hist {
			if y.Year == t.Date.Year() {
				yearExists = true
				for _, m := range y.TasksByMonth {
					if t.Date.Month() == m.Month {
						monthExists = true
						m.TaskHistory = append(m.TaskHistory, t)
						break
					}
				}

				if !monthExists {
					y.TasksByMonth = append(y.TasksByMonth, ScheduledTasksByMonth{
						Month: t.Date.Month(),
						TaskHistory: []DTOScheduledTaskListResponseData{
							t,
						},
					})
				}
				break
			}
		}

		if !yearExists {
			hist = append(hist, ScheduledTasksByYear{
				Year: t.Date.Year(),
				TasksByMonth: []ScheduledTasksByMonth{
					{
						Month: t.Date.Month(),
						TaskHistory: []DTOScheduledTaskListResponseData{
							t,
						},
					},
				},
			})
		}
	}

	return hist
}

type DTOScheduledTaskListResponseData struct {
	ScheduledTaskId string    `json:"ScheduledTaskId"`
	TaskId          string    `json:"TaskId"`
	Name            string    `json:"Name"`
	IsComplete      bool      `json:"IsComplete"`
	Date            time.Time `json:"Date"`
}

type DTOScheduledTaskNewRequest struct {
	TaskId     string    `json:"taskId"`
	IsComplete bool      `json:"isComplete"`
	Date       time.Time `json:"date"`
}

type DTOScheduledTaskUpdateRequest struct {
	ScheduledTaskId string    `json:"scheduledTaskId"`
	IsComplete      bool      `json:"isComplete"`
	Date            time.Time `json:"date"`
}

type DTOScheduledTaskCompleteRequest struct {
	ScheduledTaskId string `json:"scheduledTaskId"`
}
