package api

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type NewTaskRequest struct {
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Priority    int        `json:"priority"`
	DueDate     *time.Time `json:"dueDate"`
}

func (a *API) AddTaskHandler(c *gin.Context) {
	var newTask NewTaskRequest
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}
}

func (a *API) GetTaskByIDHandler(c *gin.Context) {}

//Gets all tasks from the database
//Query - orderby
//
func (a *API) GetAllTasksHandler(c *gin.Context) {}

//Full text search to get tasks
func (a *API) SearchTasksHandler(c *gin.Context) {}
