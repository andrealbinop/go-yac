// Package composite allows composition of config.Provider sources hierarchically.
package composite

import (
	"github.com/andrealbinop/go-yac/pkg/config"
	"strings"
)

// Provider is a config.Provider implementation to handle multiple config.Provider in
// a hierarchical fashion. If a property exists in multiple sources, the property from
// the last position source will be returned.
type Provider struct {
	// Sources is the config.Provider slice to lookup for properties.
	Sources []config.Provider
}

// String returns a string value associated with the key.
func (c *Provider) String(name string) string {
	if source := c.getFirstWithSetValue(name); source != nil {
		return source.String(name)
	}
	return config.EmptyString
}

// Int returns a int value associated with the key.
func (c *Provider) Int(name string) int {
	if source := c.getFirstWithSetValue(name); source != nil {
		return source.Int(name)
	}
	return config.Zero
}

// Float returns a float value associated with the key.
func (c *Provider) Float(name string) float64 {
	if source := c.getFirstWithSetValue(name); source != nil {
		return source.Float(name)
	}
	return config.Zero
}

// Bool returns a bool value associated with the key.
func (c *Provider) Bool(name string) bool {
	if source := c.getFirstWithSetValue(name); source != nil {
		return source.Bool(name)
	}
	return false
}

// StringSlice returns a string slice value associated with the key.
func (c *Provider) StringSlice(name string) []string {
	if source := c.getFirstWithSetValue(name); source != nil {
		return source.StringSlice(name)
	}
	return nil
}

// IsSet returns if there's a value associated with the key.
func (c *Provider) IsSet(name string) bool {
	return c.getFirstWithSetValue(name) != nil
}

// AllSettings returns all values associated with this provider as a map[string]interface{}.
func (c *Provider) AllSettings() map[string]interface{} {
	result := make(map[string]interface{})
	for _, source := range c.Sources {
		for k, v := range source.AllSettings() {
			result[k] = v
		}
	}
	return result
}

// Get returns the value associated with the key.
func (c *Provider) Get(name string) (interface{}, bool) {
	if source := c.getFirstWithSetValue(name); source != nil {
		return source.Get(name)
	}
	return nil, false
}

// Set associates a value to to the key.
func (c *Provider) Set(name string, value interface{}) interface{} {
	previous, _ := c.Get(name)
	sourceCount := len(c.Sources)
	if sourceCount > config.Zero {
		c.Sources[sourceCount-1].Set(name, value)
	}
	return previous
}

// Repository returns a merged identifier from the sources associated with this provider
func (c *Provider) Source() string {
	var sourceNames []string
	for _, source := range c.Sources {
		sourceNames = append(sourceNames, source.Source())
	}
	if len(sourceNames) == config.Zero {
		return config.EmptyString
	}
	return strings.Join(sourceNames, ">")
}

// getFirstWithSetValue returns the config.Provider that has a value associated with the key.
// The iteration in Provider.Sources is reversed, given the value lookup precedence is reversed.
func (c *Provider) getFirstWithSetValue(name string) config.Provider {
	for idx := range c.Sources {
		idx = len(c.Sources) - 1 - idx
		source := c.Sources[idx]
		if source != nil && source.IsSet(name) {
			return source
		}
	}
	return nil
}
