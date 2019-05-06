// Package provider contains default implementations from config.Provider interface
package provider

import (
	"fmt"
	"strconv"

	"github.com/andrealbinop/go-yac/pkg/config"
)

const (
	base10    = 10
	bitSize64 = 64
)

// Default implements config.Provider backed by a config.Repository
type Default struct {
	// SourceName associated with this provider
	SourceName string
	// Default containing all properties associated with this provider
	Repository config.Repository
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
		result = fmt.Sprintf("%v", value)
	}
	return result

}

// Int returns a int value associated with the key.
func (c *Default) Int(name string) int {
	var result int
	value, ok := c.Get(name)
	if ok {
		result = c.toInt(value)
	}
	return result
}

func (c *Default) toInt(value interface{}) int {
	result, ok := value.(int)
	if ok {
		return result
	}
	var resultInt64 int64
	rawValue, ok := value.(string)
	if ok {
		resultInt64, _ = strconv.ParseInt(rawValue, base10, bitSize64)
	}
	return int(resultInt64)
}

// Float returns a float value associated with the key.
func (c *Default) Float(name string) float64 {
	var result float64
	value, ok := c.Get(name)
	if ok {
		result = c.toFloat64(value)
	}
	return result
}

func (c *Default) toFloat64(value interface{}) float64 {
	result, ok := value.(float64)
	if ok {
		return result
	}
	rawValue, ok := value.(string)
	if ok {
		result, _ = strconv.ParseFloat(rawValue, bitSize64)
	}
	return result
}

// Bool returns a bool value associated with the key.
func (c *Default) Bool(name string) bool {
	var result bool
	value, ok := c.Get(name)
	if ok {
		result = c.toBool(value)
	}
	return result
}

func (c *Default) toBool(value interface{}) bool {
	result, ok := value.(bool)
	if ok {
		return result
	}
	rawValue, ok := value.(string)
	if ok {
		result, _ = strconv.ParseBool(rawValue)
	}
	return result
}

// StringSlice returns a string slice value associated with the key.
func (c *Default) StringSlice(name string) []string {
	var result []string
	value, ok := c.Get(name)
	if ok {
		result = c.toStringSlice(value)
	}
	return result
}

func (c *Default) toStringSlice(value interface{}) []string {
	var result []string
	result, ok := value.([]string)
	if ok {
		return result
	}
	rawValue, ok := value.([]interface{})
	if ok {
		for _, str := range rawValue {
			result = append(result, fmt.Sprint(str))
		}
	}
	return result
}

// IsSet returns if there's a value associated with the key.
func (c *Default) IsSet(name string) bool {
	return c.Repository.IsSet(name)
}

// AllSettings returns all values associated with this provider as a map[string]interface{}.
func (c *Default) AllSettings() map[string]interface{} {
	return c.Repository.AllSettings()
}

// Get returns the value associated with the key.
func (c *Default) Get(name string) (interface{}, bool) {
	return c.Repository.Get(name)
}

// Set associates a value to to the key.
func (c *Default) Set(name string, value interface{}) interface{} {
	return c.Repository.Set(name, value)
}
