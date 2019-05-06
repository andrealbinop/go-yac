package provider

import (
	"github.com/andrealbinop/go-yac/internal/mocks"
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	nonExistent = "non_existent"
	existent    = "existent"
	valString   = "string"
	sourceName  = "mapSource"
	valInt      = 1
	valFloat    = 1.0
	valBool     = true
)

func TestPropertyNotFoundString(t *testing.T) {
	provider, repository := providerWithProperty(nonExistent, nil)
	assert.Empty(t, provider.String(nonExistent))
	repository.AssertExpectations(t)
}

func TestPropertyNotFoundInt(t *testing.T) {
	provider, repository := providerWithProperty(nonExistent, nil)
	assert.Zero(t, provider.Int(nonExistent))
	repository.AssertExpectations(t)
}

func TestPropertyNotFoundFloat(t *testing.T) {
	provider, repository := providerWithProperty(nonExistent, nil)
	assert.Zero(t, provider.Float(nonExistent))
	repository.AssertExpectations(t)
}

func TestPropertyNotFoundBool(t *testing.T) {
	provider, repository := providerWithProperty(nonExistent, nil)
	assert.False(t, provider.Bool(nonExistent))
	repository.AssertExpectations(t)
}

func TestPropertyNotFoundStringSlice(t *testing.T) {
	provider, repository := providerWithProperty(nonExistent, nil)
	assert.Empty(t, provider.StringSlice(nonExistent))
	repository.AssertExpectations(t)
}

func TestPropertyNotFoundGet(t *testing.T) {
	provider, repository := providerWithProperty(nonExistent, nil)
	value, ok := provider.Get(nonExistent)
	assert.Nil(t, value)
	assert.False(t, ok)
	repository.AssertExpectations(t)
}

func TestPropertyNotSet(t *testing.T) {
	repository := mocks.Repository{}
	provider := Default{
		Repository: &repository,
	}
	repository.On("IsSet", nonExistent).Return(false)
	assert.False(t, provider.IsSet(nonExistent))
	repository.AssertExpectations(t)
}

func TestNoProperties(t *testing.T) {
	repository := mocks.Repository{}
	provider := Default{
		Repository: &repository,
	}
	repository.On("AllSettings").Return(nil)
	assert.Empty(t, provider.AllSettings())
	repository.AssertExpectations(t)
}

func TestSource(t *testing.T) {
	provider := Default{
		SourceName: sourceName,
	}
	assert.Equal(t, sourceName, provider.Source())
}

func TestSetProperty(t *testing.T) {
	repository := mocks.Repository{}
	provider := Default{
		Repository: &repository,
	}
	repository.On("Set", existent, valString).Return(nil)
	provider.Set(existent, valString)
	repository.AssertExpectations(t)
}

func TestPropertyGet(t *testing.T) {
	provider, repository := providerWithProperty(existent, valString)
	result, ok := provider.Get(existent)
	assert.Equal(t, valString, result)
	assert.True(t, ok)
	repository.AssertExpectations(t)
}

func TestPropertyString(t *testing.T) {
	provider, repository := providerWithProperty(existent, valString)
	assert.Equal(t, valString, provider.String(existent))
	repository.AssertExpectations(t)
}

func TestPropertyInt(t *testing.T) {
	provider, repository := providerWithProperty(existent, valInt)
	assert.Equal(t, valInt, provider.Int(existent))
	repository.AssertExpectations(t)
}

func TestPropertyIntWithRawValue(t *testing.T) {
	provider, repository := providerWithProperty(existent, "1")
	assert.Equal(t, valInt, provider.Int(existent))
	repository.AssertExpectations(t)
}

func TestPropertyFloat(t *testing.T) {
	provider, repository := providerWithProperty(existent, valFloat)
	assert.Equal(t, valFloat, provider.Float(existent))
	repository.AssertExpectations(t)
}

func TestPropertyFloatWithRawValue(t *testing.T) {
	provider, repository := providerWithProperty(existent, "1.0")
	assert.Equal(t, valFloat, provider.Float(existent))
	repository.AssertExpectations(t)
}

func TestPropertyBool(t *testing.T) {
	provider, repository := providerWithProperty(existent, valBool)
	assert.Equal(t, valBool, provider.Bool(existent))
	repository.AssertExpectations(t)
}

func TestPropertyBoolWithRawValue(t *testing.T) {
	provider, repository := providerWithProperty(existent, "true")
	assert.Equal(t, valBool, provider.Bool(existent))
	repository.AssertExpectations(t)
}

func TestPropertyStringSlice(t *testing.T) {
	value := []string{valString}
	provider, repository := providerWithProperty(existent, value)
	assert.Equal(t, value, provider.StringSlice(existent))
	repository.AssertExpectations(t)
}

func TestPropertyStringSliceWithRawValue(t *testing.T) {
	value := []interface{}{valString, valString}
	expected := []string{valString, valString}
	provider, repository := providerWithProperty(existent, value)
	assert.Equal(t, expected, provider.StringSlice(existent))
	repository.AssertExpectations(t)
}

func providerWithProperty(key string, value interface{}) (config.Provider, mocks.Repository) {
	repository := mocks.Repository{}
	repository.On("Get", key).Return(value, value != nil)
	provider := Default{
		Repository: &repository,
	}
	return &provider, repository
}
