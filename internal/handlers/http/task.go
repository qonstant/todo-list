package http

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
	_ "todo-list/docs"

	"todo-list/db/sqlc"
	"todo-list/internal/database"
	"todo-list/internal/validators"

	"github.com/go-chi/chi/v5"
)

// RegisterTaskRoutes sets up the routes for task handlers
func RegisterTaskRoutes(router *chi.Mux) {
	router.Post("/tasks", CreateTask)
	router.Put("/tasks/{id}", UpdateTask)
	router.Delete("/tasks/{id}", DeleteTask)
	router.Put("/tasks/{id}/done", MarkTaskDone)
	router.Get("/todo-list/tasks", ListTasks)
}

// CreateTask creates a new task
// @Summary Create a new task
// @Description Create a new task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body createTaskInput true "Task Input"
// @Success 201 {object} db.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Error creating task"
// @Router /api/todo-list/tasks [post]
func CreateTask(w http.ResponseWriter, r *http.Request) {
	var input createTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	activeAt, err := time.Parse("2006-01-02", input.ActiveAt)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	params := db.CreateTaskParams{
		Title:    input.Title,
		ActiveAt: activeAt,
	}

	task := db.Task{
		Title:    params.Title,
		ActiveAt: params.ActiveAt,
		Done:     input.Done,
	}

	if err := validators.ValidateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdTask, err := db.New(database.DB).CreateTask(context.Background(), params)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdTask)
}

// UpdateTask updates an existing task
// @Summary Update an existing task
// @Description Update an existing task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param id path int true "Task ID"
// @Param task body updateTaskInput true "Task Input"
// @Success 200 {object} db.Task
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Error updating task"
// @Router /api/todo-list/tasks/{id} [put]
func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var input updateTaskInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	activeAt, err := time.Parse("2006-01-02", input.ActiveAt)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	params := db.UpdateTaskParams{
		ID:       id,
		Title:    input.Title,
		ActiveAt: activeAt,
		Done:     input.Done,
	}

	task := db.Task{
		ID:       params.ID,
		Title:    params.Title,
		ActiveAt: params.ActiveAt,
		Done:     input.Done,
	}

	if err := validators.ValidateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedTask, err := db.New(database.DB).UpdateTask(context.Background(), params)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedTask)
}

// DeleteTask deletes a task
// @Summary Delete a task
// @Description Delete a task
// @Tags Tasks
// @Param id path int true "Task ID"
// @Success 204
// @Failure 400 {string} string "Invalid task ID"
// @Failure 500 {string} string "Error deleting task"
// @Router /api/todo-list/tasks/{id} [delete]
func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	err = db.New(database.DB).DeleteTask(context.Background(), id)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// MarkTaskDone marks a task as done
// @Summary Mark a task as done
// @Description Mark a task as done
// @Tags Tasks
// @Param id path int true "Task ID"
// @Success 200 {object} db.Task
// @Failure 400 {string} string "Invalid task ID"
// @Failure 500 {string} string "Error marking task as done"
// @Router /api/todo-list/tasks/{id}/done [put]
func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
    id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
    if err != nil {
        http.Error(w, "Invalid task ID", http.StatusBadRequest)
        return
    }

    existingTask, err := db.New(database.DB).GetTask(context.Background(), id)
    if err != nil {
        http.Error(w, "Error fetching task", http.StatusInternalServerError)
        return
    }

    params := db.UpdateTaskParams{
        ID:       id,
        Title:    existingTask.Title,
        ActiveAt: existingTask.ActiveAt,
        Done:     true,
    }

    if err := validators.ValidateTask(existingTask); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    updatedTask, err := db.New(database.DB).UpdateTask(context.Background(), params)
    if err != nil {
        http.Error(w, "Error marking task as done", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(updatedTask)
}


// ListTasks lists all tasks
// @Summary List all tasks
// @Description List all tasks
// @Tags Tasks
// @Produce json
// @Success 200 {array} db.Task
// @Failure 500 {string} string "Error retrieving tasks"
// @Router /api/todo-list/tasks [get]
func ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.New(database.DB).ListTasks(context.Background())
	if err != nil {
		http.Error(w, "Error retrieving tasks", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}

type createTaskInput struct {
	Title    string `json:"title"`
	ActiveAt string `json:"active_at"`
	Done     bool   `json:"done"`
}

type updateTaskInput struct {
	Title    string `json:"title"`
	ActiveAt string `json:"active_at"`
	Done     bool   `json:"done"`
}
