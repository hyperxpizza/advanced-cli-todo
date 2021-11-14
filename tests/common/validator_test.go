package main

import (
	"testing"
	"time"

	"github.com/hyperxpizza/advanced-cli-todo/internal/customErrors"
	"github.com/hyperxpizza/advanced-cli-todo/internal/models"
	"github.com/hyperxpizza/advanced-cli-todo/internal/validator"
	"github.com/stretchr/testify/assert"
)

func TestValidateNewTask(t *testing.T) {

	task := models.NewTaskRequest{
		Title:       "Create a new cli application",
		Description: "A task to create an advanced cli todo list",
		Priority:    10,
		DueDate:     time.Now().Add(time.Hour * 168),
	}

	t.Run("Validate Task OK", func(t *testing.T) {
		err := validator.ValidateNewTask(task)
		assert.NoError(t, err)
	})

	t.Run("Validate Task Empty Title", func(t *testing.T) {
		task.Title = ""
		err := validator.ValidateNewTask(task)
		assert.Error(t, err, customErrors.ErrTaskTitleNotValid)
	})

	t.Run("Validate Task Wrong Title", func(t *testing.T) {
		task.Title = "  \tabc\t "
		err := validator.ValidateNewTask(task)
		assert.Error(t, err, customErrors.ErrTaskTitleNotValid)
	})

	t.Run("Validate Task Title Too Long", func(t *testing.T) {
		task.Title = `
		Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse vel risus hendrerit, mollis ipsum eu, congue risus. Aliquam porta velit vel ante commodo consectetur eget in neque. Nulla tincidunt ante sed orci pharetra, et mattis ipsum vulputate. Quisque cursus, mauris vitae sodales cursus, lorem ipsum tincidunt nulla, in tempus turpis quam ornare elit. Vestibulum placerat nisl in enim mollis, quis sollicitudin nibh finibus. Integer porta orci tellus, sed maximus felis tempus quis. Suspendisse convallis leo mauris, vitae varius nulla aliquet sit amet. Pellentesque varius, lectus sit amet feugiat ullamcorper, sem purus blandit turpis, eget cursus est lectus id mauris. Aenean odio turpis, lobortis ac iaculis vitae, faucibus sit amet tortor. Ut sapien nulla, scelerisque ac convallis non, dignissim sed enim. Suspendisse purus neque, semper non libero vel, sagittis gravida sapien. Donec urna tellus, fringilla vitae tincidunt ac, vehicula vitae urna. Sed eu faucibus risus. 
		`
		err := validator.ValidateNewTask(task)
		assert.Error(t, err, customErrors.ErrTaskTitleNotValid)
	})

	t.Run("Validate Task Wrong Priority", func(t *testing.T) {
		task.Priority = 11
		err := validator.ValidateNewTask(task)
		assert.Error(t, err, customErrors.ErrTaskPriorityNotValid)
	})
}
