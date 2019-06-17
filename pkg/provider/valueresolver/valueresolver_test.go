package valueresolver

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/andrealbinop/go-yac/pkg/provider"
	"github.com/andrealbinop/go-yac/pkg/repository"
)

func TestValueResolverDefault_GetPropertyNotFound(t *testing.T) {
	cfg := createCfgProviderDefaultValueResolver()
	value, ok := cfg.Get("nonexistent")
	assert.Nil(t, value)
	assert.False(t, ok)
}

func TestValueResolverDefault_GetPropertyFound(t *testing.T) {
	cfg := createCfgProviderDefaultValueResolver()
	propertyExistent := "existent"
	cfg.Set(propertyExistent, propertyExistent)
	value, ok := cfg.Get(propertyExistent)
	assert.Equal(t, propertyExistent, value)
	assert.True(t, ok)
}

func TestValueResolverDefault_IsSetPropertyNotFound(t *testing.T) {
	cfg := createCfgProviderDefaultValueResolver()
	existent := cfg.IsSet("nonexistent")
	assert.False(t, existent)
}

func TestValueResolverDefault_IsSetPropertyFound(t *testing.T) {
	cfg := createCfgProviderDefaultValueResolver()
	propertyExistent := "existent"
	cfg.Set(propertyExistent, propertyExistent)
	existent := cfg.IsSet(propertyExistent)
	assert.True(t, existent)
}

func TestRelaxedName_GetPropertyNotFound(t *testing.T) {
	cfg := createCfgProviderRelaxedValueResolver()
	value, ok := cfg.Get("nonexistent")
	assert.Nil(t, value)
	assert.False(t, ok)
}

func TestRelaxedName_GetPropertyFound(t *testing.T) {
	cfg := createCfgProviderRelaxedValueResolver()
	propertyExistent := "existent"
	cfg.Set(propertyExistent, propertyExistent)
	value, ok := cfg.Get(propertyExistent)
	assert.Equal(t, propertyExistent, value)
	assert.True(t, ok)
}

func TestRelaxedName_IsSetPropertyNotFound(t *testing.T) {
	cfg := createCfgProviderRelaxedValueResolver()
	existent := cfg.IsSet("nonexistent")
	assert.False(t, existent)
}

func TestRelaxedName_IsSetPropertyFound(t *testing.T) {
	cfg := createCfgProviderRelaxedValueResolver()
	propertyExistent := "existent"
	cfg.Set(propertyExistent, propertyExistent)
	existent := cfg.IsSet(propertyExistent)
	assert.True(t, existent)
}

func createCfgProviderRelaxedValueResolver() provider.Default {
	cfg := createCfgProviderDefaultValueResolver()
	cfg.ValueResolver = &valueResolverRelaxedName{}
	return cfg
}

func TestRelaxedName_GetPropertyFoundEnvFormat(t *testing.T) {
	cfg := createCfgProviderRelaxedValueResolver()
	valueExistent := "value"
	cfg.Set("prop.existent", valueExistent)
	value, ok := cfg.Get("PROP_EXISTENT")
	assert.Equal(t, valueExistent, value)
	assert.True(t, ok)
}

func TestRelaxedName_GetPropertyFoundYmlFormat(t *testing.T) {
	cfg := createCfgProviderRelaxedValueResolver()
	valueExistent := "value"
	cfg.Set("PROP_EXISTENT", valueExistent)
	value, ok := cfg.Get("prop.existent")
	assert.Equal(t, valueExistent, value)
	assert.True(t, ok)
}

func TestRelaxedName_IsSetPropertyFoundEnvFormat(t *testing.T) {
	cfg := createCfgProviderRelaxedValueResolver()
	cfg.Set("prop.existent", "value")
	ok := cfg.IsSet("PROP_EXISTENT")
	assert.True(t, ok)
}

func TestRelaxedName_IsSetPropertyFoundYmlFormat(t *testing.T) {
	cfg := createCfgProviderRelaxedValueResolver()
	cfg.Set("PROP_EXISTENT", "value")
	ok := cfg.IsSet("prop.existent")
	assert.True(t, ok)
}

func createCfgProviderDefaultValueResolver() provider.Default {
	return provider.Default{
		Repository:    &repository.Map{},
		ValueResolver: &Default{},
	}
}

type valueResolverRelaxedName struct {
}

func (v *valueResolverRelaxedName) IsSet(name string, repository config.Repository) bool {
	for _, name := range v.parseRelaxedNames(name) {
		if ok := repository.IsSet(name); ok {
			return true
		}
	}
	return false
}

func (v *valueResolverRelaxedName) Resolve(name string, repository config.Repository) (interface{}, bool) {
	for _, name := range v.parseRelaxedNames(name) {
		if v, ok := repository.Get(name); ok {
			return v, true
		}
	}
	return nil, false
}

func (v *valueResolverRelaxedName) parseRelaxedNames(name string) []string {
	const (
		propSeparator = "."
		envSeparator  = "_"
	)
	names := []string{name}
	nameUpperEnv := strings.ToUpper(strings.Replace(name, propSeparator, envSeparator, -1))
	if name != nameUpperEnv {
		names = append(names, nameUpperEnv)
	}
	nameLowerProp := strings.ToLower(strings.Replace(nameUpperEnv, envSeparator, propSeparator, -1))
	if name != nameLowerProp {
		names = append(names, nameLowerProp)
	}
	return names
}
