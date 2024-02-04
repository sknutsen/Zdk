package models

import (
	"github.com/sknutsen/Zdk/constants"
)

type ClientStateBase struct {
	PageId      uint
	LoggedIn    bool
	UserProfile UserProfile
}

func getBaseState(profile interface{}) ClientStateBase {
	state := ClientStateBase{}

	if profile != nil {
		state.LoggedIn = true
		state.UserProfile = GetUserProfile(profile.(map[string]interface{}))
	} else {
		state.LoggedIn = false
	}

	return state
}

type IndexState struct {
	State ClientStateBase
}

func GetIndexState(profile interface{}) IndexState {
	state := getBaseState(profile)

	state.PageId = constants.IndexPageId

	return IndexState{
		State: state,
	}
}

type ShoppingListsState struct {
	State ClientStateBase
}

func GetShoppingListsState(profile interface{}) ShoppingListsState {
	state := getBaseState(profile)

	state.PageId = constants.ShoppingListsPageId

	return ShoppingListsState{
		State: state,
	}
}

type TasksState struct {
	State                ClientStateBase
	CategoryContext      *DTOCategoryUpdateRequest
	Categories           []Category
	CurrentScheduledItem *DTOScheduledTaskListResponseData
	TaskHistory          []ScheduledTasksByYear
	TaskContext          *DTOTaskUpdateRequest
	Tasks                []Task
}

func GetTasksState(profile interface{}) TasksState {
	state := getBaseState(profile)

	state.PageId = constants.TasksPageId

	return TasksState{
		State: state,
	}
}

type TodoState struct {
	State ClientStateBase
}

func GetTdodState(profile interface{}) TodoState {
	state := getBaseState(profile)

	state.PageId = constants.TasksPageId

	return TodoState{
		State: state,
	}
}

type UserState struct {
	State ClientStateBase
}

func GetUserState(profile interface{}) UserState {
	state := getBaseState(profile)

	return UserState{
		State: state,
	}
}

type WorkoutState struct {
	State          ClientStateBase
	WorkoutContext *DTOWorkoutUpdateRequest
	Workouts       []Workout
}

func GetWorkoutState(profile interface{}) WorkoutState {
	state := getBaseState(profile)

	state.PageId = constants.WorkoutPageId

	return WorkoutState{
		State: state,
	}
}
