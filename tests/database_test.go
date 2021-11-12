package main

import (
	"testing"

	"github.com/hyperxpizza/advanced-cli-todo/internal/common"
	"github.com/hyperxpizza/advanced-cli-todo/internal/config"
	"github.com/hyperxpizza/advanced-cli-todo/internal/db"
	"github.com/stretchr/testify/assert"
)

func TestConnectToDB(t *testing.T) {
	if configPathPtr == nil || *configPathPtr == "" {
		t.Log("config flag has to be set! usage: --config=/path/to/config")
		t.Fail()
		return
	}

	c, err := config.NewConfig(*configPathPtr)
	assert.NoError(t, err)

	logger := common.NewLogger(*loglevel)
	database, err := db.NewDatabase(c, logger)
	assert.NoError(t, err)

	defer database.Close()
}
