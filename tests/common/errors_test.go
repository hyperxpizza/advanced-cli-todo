package main

import (
	"testing"

	"github.com/hyperxpizza/advanced-cli-todo/internal/customErrors"
	"github.com/stretchr/testify/assert"
)

func TestErrorsWrapper(t *testing.T) {
	err1 := customErrors.Wrap(customErrors.ErrFileNotFound)
	assert.Error(t, err1)
	assert.Equal(t, err1.Error(), customErrors.ErrFileNotFound)
}
