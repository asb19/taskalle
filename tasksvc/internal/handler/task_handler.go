package handler

import (
	"encoding/json"
	"net/http"

	"github.com/asb19/tasksvc/internal/model"
	"github.com/asb19/tasksvc/internal/service"
	"github.com/asb19/tasksvc/internal/utils"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Handler struct {
	service service.TaskService
}

func NewHandler(service service.TaskService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateTask(w http.ResponseWriter, r *http.Request) {
	var req struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		Status      string `json:"status"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task, err := h.service.CreateTask(r.Context(), &model.Task{Title: req.Title, Description: req.Description, Status: model.TaskStatus(req.Status)})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

// ListTasks godoc
// @Summary List tasks
// @Description Get all tasks with optional pagination and filtering by status
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param status query string false "Task status filter"
// @Param limit query int false "Limit number of tasks" default(10)
// @Param page query int false "Skip tasks for pagination" default(1)
// @Success 200 {array} model.Task
// @Failure 500 {object} map[string]string
// @Router /tasks [get]
func (h *Handler) GetAllTask(w http.ResponseWriter, r *http.Request) {

	q := r.URL.Query()
	status := q.Get("status")
	limit := utils.ParseQueryParamInt(q, "limit", 10) // default = 10
	page := utils.ParseQueryParamInt(q, "page", 1)    // default = 0
	tasks, err := h.service.GetTasks(r.Context(), status, page, limit)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

// GetTask godoc
// @Summary Get a task by ID
// @Description Get a single task including assigned user info
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path string true "Task ID (UUID)"
// @Success 200 {object} model.Task
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [get]
func (h *Handler) GetTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := uuid.Parse(idStr)
	t, err := h.service.GetTaskByID(r.Context(), id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(t)
}

// UpdateTask godoc
// @Summary Update a task by ID
// @Description Update a task's title, description, status, or assigned user
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID (UUID)"
// @Param task body model.Task true "Task object"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /tasks/{id} [put]
func (h *Handler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := uuid.Parse(idStr)
	var t model.Task
	if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	t.Id = id
	if err := h.service.UpdateTask(r.Context(), &t); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "updated"})
}

// DeleteTask godoc
// @Summary Delete a task by ID
// @Description Delete a task from the database by its ID
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID (UUID)"
// @Success 200 {object} map[string]string "Task deleted successfully"
// @Failure 400 {object} map[string]string "Invalid task ID"
// @Failure 404 {object} map[string]string "Task not found"
// @Router /tasks/{id} [delete]
func (h *Handler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	idStr := mux.Vars(r)["id"]
	id, _ := uuid.Parse(idStr)
	if err := h.service.DeleteTask(r.Context(), id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
}
