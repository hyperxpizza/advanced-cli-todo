package main

import (
	"flag"
	"testing"

	"github.com/hyperxpizza/advanced-cli-todo/internal/common"
	"github.com/stretchr/testify/assert"
)

func TestCheckIfFileExists(t *testing.T) {
	flag.Parse()

	if configPathPtr == nil || *configPathPtr == "" {
		t.Log("config flag has to be set! usage: --config=/path/to/config")
		t.Fail()
		return
	}

	err := common.CheckIfFileExists(*configPathPtr)
	assert.NoError(t, err)
}

func TestReadDatabaseSchema(t *testing.T) {
	flag.Parse()

	if filePtr == nil || *filePtr == "" {
		t.Log("path flag has to be set! usage: --file=/path")
		t.Fail()
		return
	}

	if !common.CheckFileExtension(*filePtr, ".sql") {
		t.Log("File needs a .sql extenstion!")
		t.Fail()
		return
	}

	data, err := common.ReadFile(*filePtr)
	assert.NoError(t, err)

	toCheck := "create table tasks ( id serial primary key unique not null, title varchar(100) not null, description text, done boolean not null, priority integer not null, dueDate timestamp, created timestamp not null, updated timestamp not null);"
	assert.Equal(t, toCheck, string(data))

}
