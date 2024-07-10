package api

// import (
// 	"net/http"
// 	"strconv"
// 	"time"

// 	db "Simple-Bank/db/sqlc"

// 	"github.com/gin-gonic/gin"
// )

// type Server struct {
// 	store  db.Store
// 	router *gin.Engine
// }

// func NewServer(store db.Store) *Server {
// 	server := &Server{
// 		store: store,
// 	}
// 	server.setupRouter()
// 	return server
// }

// func errorResponse(c *gin.Context, code int, message string) {
// 	c.JSON(code, gin.H{"error": message})
// }

// func (server *Server) Start(address string) error {
// 	return server.router.Run(address)
// }

// func (server *Server) setupRouter() {
// 	router := gin.Default()

// 	router.POST("/tasks", server.createTask)
// 	router.GET("/tasks/:id", server.getTask)
// 	router.GET("/tasks", server.listTasks)
// 	router.PUT("/tasks/:id", server.updateTask)
// 	router.DELETE("/tasks/:id", server.deleteTask)

// 	server.router = router
// }

// func (server *Server) createTask(c *gin.Context) {
// 	var request struct {
// 		Title    string    `json:"title"`
// 		ActiveAt time.Time `json:"active_at"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		errorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	task, err := server.store.CreateTask(c.Request.Context(), request.Title, request.ActiveAt)
// 	if err != nil {
// 		errorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, task)
// }

// func (server *Server) getTask(c *gin.Context) {
// 	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		errorResponse(c, http.StatusBadRequest, "invalid task ID")
// 		return
// 	}

// 	task, err := server.store.GetTask(c.Request.Context(), taskID)
// 	if err != nil {
// 		errorResponse(c, http.StatusNotFound, "task not found")
// 		return
// 	}

// 	c.JSON(http.StatusOK, task)
// }

// func (server *Server) listTasks(c *gin.Context) {
// 	tasks, err := server.store.ListTasks(c.Request.Context())
// 	if err != nil {
// 		errorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, tasks)
// }

// func (server *Server) updateTask(c *gin.Context) {
// 	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		errorResponse(c, http.StatusBadRequest, "invalid task ID")
// 		return
// 	}

// 	var request struct {
// 		Title    string    `json:"title"`
// 		ActiveAt time.Time `json:"active_at"`
// 	}

// 	if err := c.ShouldBindJSON(&request); err != nil {
// 		errorResponse(c, http.StatusBadRequest, err.Error())
// 		return
// 	}

// 	task, err := server.store.UpdateTask(c.Request.Context(), taskID, request.Title, request.ActiveAt)
// 	if err != nil {
// 		errorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.JSON(http.StatusOK, task)
// }

// func (server *Server) deleteTask(c *gin.Context) {
// 	taskID, err := strconv.ParseInt(c.Param("id"), 10, 64)
// 	if err != nil {
// 		errorResponse(c, http.StatusBadRequest, "invalid task ID")
// 		return
// 	}

// 	err = server.store.DeleteTask(c.Request.Context(), taskID)
// 	if err != nil {
// 		errorResponse(c, http.StatusInternalServerError, err.Error())
// 		return
// 	}

// 	c.Status(http.StatusNoContent)
// }
