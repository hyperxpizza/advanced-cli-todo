package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
)

func (a *API) setupAndRunRouter(c *config.Config) {
	router := gin.Default()
	router.Use(a.corsMiddleware())
	addr := fmt.Sprintf("%s:%d", c.API.Host, c.API.Port)

	router.Run(addr)
}
