// Package provider contains default implementations from config.Provider interface
package provider

import (
	"github.com/andrealbinop/go-yac/pkg/config"
)

// Default implements config.Provider backed by a config.Repository
type Default struct {
	// SourceName associated with this provider
	SourceName string
	// Default containing all properties associated with this provider
	Repository config.Repository
	// ValueResolver to retrieve and verify the existence of values in repository
	ValueResolver config.ValueResolver
	// ValueConverter is used to convert values providing from repository
	ValueConverter config.ValueConverter
}

// Source returns the source name associated with this provider
func (c *Default) Source() string {
	return c.SourceName
}

// String returns a string value associated with the key.
func (c *Default) String(name string) string {
	var result string
	value, ok := c.Get(name)
	if ok {
		result = c.ValueConverter.ToString(value)
	}
	return result

}

// Int returns a int value associated with the key.
func (c *Default) Int(name string) int {
	var result int
	value, ok := c.Get(name)
	if ok {
		result = c.ValueConverter.ToInt(value)
	}
	return result
}

// Float returns a float value associated with the key.
func (c *Default) Float(name string) float64 {
	var result float64
	value, ok := c.Get(name)
	if ok {
		result = c.ValueConverter.ToFloat(value)
	}
	return result
}

// Bool returns a bool value associated with the key.
func (c *Default) Bool(name string) bool {
	var result bool
	value, ok := c.Get(name)
	if ok {
		result = c.ValueConverter.ToBool(value)
	}
	return result
}

// StringSlice returns a string slice value associated with the key.
func (c *Default) StringSlice(name string) []string {
	var result []string
	value, ok := c.Get(name)
	if ok {
		result = c.ValueConverter.ToStringSlice(value)
	}
	return result
}


// IntSlice returns a int slice value associated with the key.
func (c *Default) IntSlice(name string) []int {
	var result []int
	value, ok := c.Get(name)
	if ok {
		result = c.ValueConverter.ToIntSlice(value)
	}
	return result
}

// IsSet returns if there's a value associated with the key.
func (c *Default) IsSet(name string) bool {
	return c.ValueResolver.IsSet(name, c.Repository)
}

// AllSettings returns all values associated with this provider as a map[string]interface{}.
func (c *Default) AllSettings() map[string]interface{} {
	return c.Repository.AllSettings()
}

// Get returns the value associated with the key.
func (c *Default) Get(name string) (interface{}, bool) {
	return c.ValueResolver.Resolve(name, c.Repository)
}

// Set associates a value to to the key.
func (c *Default) Set(name string, value interface{}) interface{} {
	return c.Repository.Set(name, value)
}
