package handlers

import (
	"context"
	"net/http"

	taskservice "github.com/TheEastWantsThis/OldNew/internal/taskService"
	"github.com/TheEastWantsThis/OldNew/internal/web/tasks"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskservice.MainTaskService
}

func NewTaskHandler(s taskservice.MainTaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetTasks(ctx context.Context, request tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAllTask()
	if err != nil {
		return nil, err
	}

	var response tasks.GetTasks200JSONResponse

	for _, t := range allTasks {
		id := uint(t.ID)
		task := t.Task
		accomplished := t.Accomplishment

		response = append(response, tasks.Task{
			Id:             &id,
			Task:           &task,
			Accomplishment: &accomplished,
		})
	}

	return response, nil
}

func (h *TaskHandler) PostTasks(ctx context.Context, request tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	taskRequest := request.Body

	taskToCreate := taskservice.RequestBodyTask{
		Task:           *taskRequest.Task,
		Accomplishment: *taskRequest.Accomplishment,
	}
	createdTask, err := h.service.CreateTask(taskToCreate)

	if err != nil {
		return nil, err
	}

	id := uint(createdTask.ID)

	response := tasks.PostTasks201JSONResponse{
		Id:             &id,
		Task:           &createdTask.Task,
		Accomplishment: &createdTask.Accomplishment,
	}

	return response, nil
}
func (h *TaskHandler) PatchTasksId(ctx context.Context, request tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {

	update := request.Body

	id := request.Id

	updated, err := h.service.UpdateTask(int(id), taskservice.RequestBodyTask{
		Task:           *update.Task,
		Accomplishment: *update.Accomplishment,
	})

	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Could not update task")
	}

	idUint := uint(updated.ID)
	result := tasks.Task{
		Id:             &idUint,
		Task:           &updated.Task,
		Accomplishment: &updated.Accomplishment,
	}

	return tasks.PatchTasksId201JSONResponse(result), nil
}

func (h *TaskHandler) DeleteTasksId(ctx context.Context, request tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := request.Id

	if err := h.service.DeleteTask(int(id)); err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, "Could not delete task")
	}

	return tasks.DeleteTasksId204JSONResponse{}, nil
}
