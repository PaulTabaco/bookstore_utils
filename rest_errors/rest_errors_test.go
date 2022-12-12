package rest_errors

import (
	"errors"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewError(t *testing.T) {
	err := NewError("new error message")

	assert.NotNil(t, err)
	assert.EqualValues(t, "new error message", err.Error())
}

func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("this is a message", errors.New("incorrect request"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusBadRequest, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "bad_request", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "incorrect request", err.Causes[0])
}

func TestNewNotFoundError(t *testing.T) {
	err := NewNotFoundError("this is a message", errors.New("item not found"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "not_found", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "item not found", err.Causes[0])
}

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("this is a message", errors.New("datebase error"))

	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "this is a message", err.Message)
	assert.EqualValues(t, "internal_server_error", err.Error)

	assert.NotNil(t, err.Causes)
	assert.EqualValues(t, 1, len(err.Causes))
	assert.EqualValues(t, "datebase error", err.Causes[0])
	// errBytes, _ := json.Marshal(err)
	// fmt.Println(string(errBytes))
}
