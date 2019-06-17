package provider

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andrealbinop/go-yac/internal/mocks"
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/andrealbinop/go-yac/pkg/provider/valueresolver"
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
		Repository:    &repository,
		ValueResolver: &valueresolver.Default{},
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
	parser := &mocks.ValueConverter{}
	parser.On("ToString", valString).Return(valString)
	provider, repository := providerWithPropertyAndParser(existent, valString, parser)
	assert.Equal(t, valString, provider.String(existent))
	repository.AssertExpectations(t)
	parser.AssertExpectations(t)
}

func TestPropertyInt(t *testing.T) {
	parser := &mocks.ValueConverter{}
	parser.On("ToInt", valInt).Return(valInt)
	provider, repository := providerWithPropertyAndParser(existent, valInt, parser)
	assert.Equal(t, valInt, provider.Int(existent))
	repository.AssertExpectations(t)
	parser.AssertExpectations(t)
}

func TestPropertyFloat(t *testing.T) {
	parser := &mocks.ValueConverter{}
	parser.On("ToFloat", valFloat).Return(valFloat)
	provider, repository := providerWithPropertyAndParser(existent, valFloat, parser)
	assert.Equal(t, valFloat, provider.Float(existent))
	repository.AssertExpectations(t)
	parser.AssertExpectations(t)
}

func TestPropertyBool(t *testing.T) {
	parser := &mocks.ValueConverter{}
	parser.On("ToBool", valBool).Return(valBool)
	provider, repository := providerWithPropertyAndParser(existent, valBool, parser)
	assert.Equal(t, valBool, provider.Bool(existent))
	repository.AssertExpectations(t)
	parser.AssertExpectations(t)
}

func TestPropertyStringSlice(t *testing.T) {
	parser := &mocks.ValueConverter{}
	value := []string{valString}
	parser.On("ToStringSlice", value).Return(value)
	provider, repository := providerWithPropertyAndParser(existent, value, parser)
	assert.Equal(t, value, provider.StringSlice(existent))
	repository.AssertExpectations(t)
}

func providerWithProperty(key string, value interface{}) (config.Provider, mocks.Repository) {
	return buildDefaultProviderWithProperty(key, value)
}

func providerWithPropertyAndParser(key string, value interface{}, parser *mocks.ValueConverter) (config.Provider, mocks.Repository) {
	provider, repository := buildDefaultProviderWithProperty(key, value)
	provider.ValueConverter = parser
	return provider, repository
}

func buildDefaultProviderWithProperty(key string, value interface{}) (*Default, mocks.Repository) {
	repository := mocks.Repository{}
	repository.On("Get", key).Return(value, value != nil)
	provider := Default{
		Repository:    &repository,
		ValueResolver: &valueresolver.Default{},
	}
	return &provider, repository
}
