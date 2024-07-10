package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"todo-list/internal/database"
	"todo-list/internal/models"
	"todo-list/internal/validators"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func RegisterTaskRoutes(router *chi.Mux) {
	router.Post("/api/todo-list/tasks", CreateTask)
	router.Put("/api/todo-list/tasks/{id}", UpdateTask)
	router.Delete("/api/todo-list/tasks/{id}", DeleteTask)
	router.Put("/api/todo-list/tasks/{id}/done", MarkTaskDone)
	router.Get("/api/todo-list/tasks", ListTasks)
}

func generateUUID() string {
	return uuid.New().String()
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := validators.ValidateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	task.ID = generateUUID()
	task.CreatedAt = time.Now()
	task.UpdatedAt = time.Now()

	query := `INSERT INTO tasks (id, title, active_at, done, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`
	_, err := database.DB.Exec(query, task.ID, task.Title, task.ActiveAt, task.Done, task.CreatedAt, task.UpdatedAt)
	if err != nil {
		http.Error(w, "Error creating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": task.ID})
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if err := validators.ValidateTask(task); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := `UPDATE tasks SET title = $1, active_at = $2, updated_at = $3 WHERE id = $4`
	_, err := database.DB.Exec(query, task.Title, task.ActiveAt, time.Now(), id)
	if err != nil {
		http.Error(w, "Error updating task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query := `DELETE FROM tasks WHERE id = $1`
	_, err := database.DB.Exec(query, id)
	if err != nil {
		http.Error(w, "Error deleting task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func MarkTaskDone(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	query := `UPDATE tasks SET done = $1, updated_at = $2 WHERE id = $3`
	_, err := database.DB.Exec(query, true, time.Now(), id)
	if err != nil {
		http.Error(w, "Error marking task as done", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func ListTasks(w http.ResponseWriter, r *http.Request) {
	status := r.URL.Query().Get("status")
	if status == "" {
		status = "active"
	}

	var query string
	if status == "active" {
		query = `SELECT id, title, active_at, done, created_at, updated_at FROM tasks WHERE active_at <= now() AND done = false ORDER BY created_at`
	} else {
		query = `SELECT id, title, active_at, done, created_at, updated_at FROM tasks WHERE done = true ORDER BY created_at`
	}

	rows, err := database.DB.Query(query)
	if err != nil {
		http.Error(w, "Error retrieving tasks", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	tasks := []models.Task{}
	for rows.Next() {
		var task models.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.ActiveAt, &task.Done, &task.CreatedAt, &task.UpdatedAt); err != nil {
			http.Error(w, "Error scanning task", http.StatusInternalServerError)
			return
		}
		if task.ActiveAt.Weekday() == time.Saturday || task.ActiveAt.Weekday() == time.Sunday {
			task.Title = "ВЫХОДНОЙ - " + task.Title
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "Error iterating over tasks", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(tasks)
}
