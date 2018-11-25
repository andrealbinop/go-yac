// Package repository with a map[string]interface{} backed config.Repository implementation
package repository

// Map implements config.Provider with a map[string]interface{} database
type Map struct {
	// Default containing all properties associated with this provider
	Database map[string]interface{}
}

// IsSet returns if there's a value associated with the key.
func (c *Map) IsSet(name string) bool {
	_, isSet := c.Get(name)
	return isSet
}

// AllSettings returns all values associated with this provider as a map[string]interface{}.
func (c *Map) AllSettings() map[string]interface{} {
	return c.Database
}

// Get returns the value associated with the key.
func (c *Map) Get(name string) (interface{}, bool) {
	data, present := c.Database[name]
	return data, present
}

// Set associates a value to to the key.
func (c *Map) Set(name string, value interface{}) interface{} {
	if c.Database == nil {
		c.Database = make(map[string]interface{})
	}
	previous := c.Database[name]
	c.Database[name] = value
	return previous
}
