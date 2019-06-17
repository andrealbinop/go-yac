// Package config holds main interfaces and constants to interact with configuration provisioning
package config

const (
	// EmptyString constant to be used along implementations
	EmptyString = ""
	// Zero number constant to be used along implementations
	Zero = 0
)

// Source interface for implementors to provide an identifier regarding loaded configuration
type Source interface {
	// Repository function, provides an identifier regarding loaded configuration
	Source() string
}

// Repository interface that gets and sets configuration related data
type Repository interface {
	// Set associates a value to to the key.
	Set(string, interface{}) interface{}
	// Get returns the value associated with the key and a boolean value indicating if it was found.
	Get(string) (interface{}, bool)
	// IsSet returns if there's a value associated with the key.
	IsSet(string) bool
	// AllSettings returns all values associated as a map[string]interface{}.
	AllSettings() map[string]interface{}
}

// Provider interface for configuration related data
type Provider interface {
	Source
	Repository
	// String returns a string value associated with the key.
	String(string) string
	// Int returns a int value associated with the key.
	Int(string) int
	// Float returns a float value associated with the key.
	Float(string) float64
	// Bool returns a bool value associated with the key.
	Bool(string) bool
	// StringSlice returns a string slice value associated with the key.
	StringSlice(string) []string
}

// ValueConverter interface for parser values into specific types
type ValueConverter interface {
	// ToString parse a value into string
	ToString(interface{}) string
	// ToInt parses a value into int
	ToInt(interface{}) int
	// ToFloat parses a value into float64
	ToFloat(interface{}) float64
	// ToBool parses a value into bool
	ToBool(interface{}) bool
	// ToStringSlice parses a value into string slice
	ToStringSlice(interface{}) []string
}

// ValueResolver interface to customize property recovery by name
type ValueResolver interface {
	// Resolve returns a value associated with the name.
	Resolve(string, Repository) (interface{}, bool)
	// IsSet returns if there's a value associated with the name.
	IsSet(string, Repository) bool
}

// Loader interface to represent loading config.Provider
type Loader interface {
	Source
	// Load the config.Provider
	Load() (Provider, error)
}
