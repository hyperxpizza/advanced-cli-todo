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

	//validate the new task
	if err := validator.ValidateNewTask(newTask); err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{
			"message": err.Error(),
		})
		return
	}

	//insert it into the database
	id, err := a.db.InsertTask(newTask)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
	return
}

func (a *API) GetTaskByIDHandler(c *gin.Context) {}

//Gets all tasks from the database
//Query - orderby
//
func (a *API) GetAllTasksHandler(c *gin.Context) {}

//Full text search to get tasks
//Query ?q
func (a *API) SearchTasksHandler(c *gin.Context) {}
