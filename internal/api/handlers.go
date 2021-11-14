package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/advanced-cli-todo/internal/models"
	"github.com/hyperxpizza/advanced-cli-todo/internal/validator"
)

//Inserts a new task into the database
func (a *API) AddTaskHandler(c *gin.Context) {
	var newTask models.NewTaskRequest
	//unmarshal json into struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	//validate
	if err := validator.ValidateNewTask(); err != nil {
		c.Status()
	}
}

func (a *API) GetTaskByIDHandler(c *gin.Context) {}

//Gets all tasks from the database
//Query - orderby
//
func (a *API) GetAllTasksHandler(c *gin.Context) {}

//Full text search to get tasks
//Query ?q
func (a *API) SearchTasksHandler(c *gin.Context) {}
