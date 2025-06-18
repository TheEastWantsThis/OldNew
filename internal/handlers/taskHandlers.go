package handlers

import (
	taskservice "dacode/OldNewProdleckt/internal/taskService"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	service taskservice.MainTaskService
}

func NewTaskHandler(s taskservice.MainTaskService) *TaskHandler {
	return &TaskHandler{service: s}
}

func (h *TaskHandler) GetHandler(c echo.Context) error {

	task, err := h.service.GetAllTask()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not get tasks"})
	}

	return c.JSON(http.StatusOK, task)
}

func (h *TaskHandler) PostHandler(c echo.Context) error {
	var req taskservice.RequestBodyTask
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	postTask, err := h.service.CreateTask(req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Could not create task"})
	}

	return c.JSON(http.StatusCreated, postTask)

}

func (h *TaskHandler) PatchHandler(c echo.Context) error {

	idStr := c.Param("ID")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var req taskservice.RequestBodyTask

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	updateTask, err := h.service.UpdateTask(idInt, req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Could not update task")
	}

	return c.JSON(http.StatusOK, updateTask)
}

func (h *TaskHandler) DeleteHandler(c echo.Context) error {

	idStr := c.Param("ID")

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID format")
	}

	var req taskservice.RequestBodyTask
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	if err := h.service.DeleteTask(idInt); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Could not delete task"})
	}

	return c.NoContent(http.StatusNoContent)
}
