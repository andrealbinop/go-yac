package repository

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	nonExistent = "non_existent"
	existent    = "existent"
	valString   = "string"
)

func TestPropertyNotFoundGet(t *testing.T) {
	repository := Map{}
	value, ok := repository.Get(nonExistent)
	assert.Nil(t, value)
	assert.False(t, ok)
}

func TestPropertyNotSet(t *testing.T) {
	repository := Map{}
	assert.False(t, repository.IsSet(nonExistent))
}

func TestNoProperties(t *testing.T) {
	repository := Map{}
	assert.Empty(t, repository.AllSettings())
}

func TestSetProperty(t *testing.T) {
	repository := Map{}
	assert.Empty(t, repository.AllSettings())
	assert.False(t, repository.IsSet(existent))
	value, ok := repository.Get(existent)
	assert.Nil(t, value)
	assert.False(t, ok)
	assert.Nil(t, repository.Set(existent, valString))
	assert.Len(t, repository.AllSettings(), 1)
	assert.True(t, repository.IsSet(existent))
	value, ok = repository.Get(existent)
	assert.Equal(t, valString, value)
	assert.True(t, ok)
}

func TestPropertyGet(t *testing.T) {
	repository := Map{
		Database: map[string]interface{}{
			existent: valString,
		},
	}
	value, ok := repository.Get(existent)
	assert.Equal(t, valString, value)
	assert.True(t, ok)
}
