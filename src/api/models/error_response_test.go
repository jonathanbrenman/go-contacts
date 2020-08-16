package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHttpError(t *testing.T) {
	assert := assert.New(t)
	httpError := NewHttpError(500, "internal error")
	assert.Equal(500, httpError.Code)
	assert.Equal("internal error", httpError.Error)
}