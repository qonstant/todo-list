package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"todo-list/db/sqlc"
	"todo-list/internal/database"
	"todo-list/internal/validators"

	"github.com/go-chi/chi/v5"
)

func RegisterTaskRoutes(router *chi.Mux) {
	router.Post("/api/todo-list/tasks", CreateTask)
	router.Put("/api/todo-list/tasks/{id}", UpdateTask)
	router.Delete("/api/todo-list/tasks/{id}", DeleteTask)
	router.Put("/api/todo-list/tasks/{id}/done", MarkTaskDone)
	router.Get("/api/todo-list/tasks", ListTasks)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title    string `json:"title"`
		ActiveAt string `json:"active_at"`
		Done     bool   `json:"done"`
	}
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

	// Convert to db.Task for validation
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

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	var input struct {
		Title    string `json:"title"`
		ActiveAt string `json:"active_at"`
		Done     bool   `json:"done"`
	}
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
	}

	// Convert to db.Task for validation
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

func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		http.Error(w, "Invalid task ID", http.StatusBadRequest)
		return
	}

	params := db.UpdateTaskParams{
		ID:       id,
		ActiveAt: time.Now(),
	}

	// Fetching the existing task to get its current title
	existingTask, err := db.New(database.DB).GetTask(context.Background(), id)
	if err != nil {
		http.Error(w, "Error fetching task", http.StatusInternalServerError)
		return
	}
	params.Title = existingTask.Title

	// Update done field directly on the existing task
	existingTask.Done = true

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

func ListTasks(w http.ResponseWriter, r *http.Request) {
	tasks, err := db.New(database.DB).ListTasks(context.Background())
	if err != nil {
		http.Error(w, "Error retrieving tasks", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
