package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
)

func (a *API) setupAndRunRouter(c *config.Config) {
	router := gin.Default()
	//gin.SetMode(c.API.Mode)
	//use corse middleware
	router.Use(a.corsMiddleware())

	addr := fmt.Sprintf("%s:%d", c.API.Host, c.API.Port)

	tasks := router.Group("/api/v1/tasks")
	{
		tasks.POST("", a.AddTaskHandler)
		tasks.GET("", a.GetAllTasksHandler)
		tasks.GET("/:id", a.GetTaskByIDHandler)
	}

	search := router.Group("/api/v1/search")
	{
		search.GET("", a.SearchTasksHandler)
	}

	router.Run(addr)
}
