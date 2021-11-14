package api

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/advanced-cli-todo/internal/models"
	"github.com/hyperxpizza/advanced-cli-todo/internal/validator"
)

//Inserts a new task into the database
func (a *API) AddTaskHandler(c *gin.Context) {
	var newTask models.NewTaskRequest
	//unmarshal json into struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
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
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"id": id,
	})
}

func (a *API) GetTaskByIDHandler(c *gin.Context) {}

//Gets all tasks from the database
//Query - orderby
//
func (a *API) GetAllTasksHandler(c *gin.Context) {}

//Full text search to get tasks
//Query ?q
func (a *API) SearchTasksHandler(c *gin.Context) {}

//Updates the done state of the task with provided id
func (a *API) UpdateDoneHandler(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	doneQuery := c.Query("done")
	var done bool
	switch doneQuery {
	case "true":
		done = true
	case "false":
		done = false
	default:
		c.Status(http.StatusBadRequest)
		return
	}

	err = a.db.UpdateDone(id, done)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.Status(http.StatusNotFound)
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}
