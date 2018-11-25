package main

import "github.com/spf13/viper"

// ViperProvider is an implementation for config.Provider
type ViperProvider struct {
	// Viper instance associated with this provider
	Viper *viper.Viper
}

// Source returns associated viper configuration source
func (v *ViperProvider) Source() string {
	return v.Viper.ConfigFileUsed()
}

// Set delegates data to viper.Set
func (v *ViperProvider) Set(name string, value interface{}) interface{} {
	previous, _ := v.Get(name)
	viper.Set(name, value)
	return previous
}

// Get configuration data from viper
func (v *ViperProvider) Get(name string) (interface{}, bool) {
	return v.Viper.Get(name), v.IsSet(name)
}

// IsSet checks if configuration exists in viper
func (v *ViperProvider) IsSet(name string) bool {
	return v.Viper.IsSet(name)
}

// AllSettings delegates to viper.AllSettings
func (v *ViperProvider) AllSettings() map[string]interface{} {
	return v.Viper.AllSettings()
}

// String delegates to viper.GetString
func (v *ViperProvider) String(name string) string {
	return v.Viper.GetString(name)
}

// Int delegates to viper.GetInt
func (v *ViperProvider) Int(name string) int {
	return v.Viper.GetInt(name)
}

// Float delegates to viper.GetFloat64
func (v *ViperProvider) Float(name string) float64 {
	return v.Viper.GetFloat64(name)
}

// Bool delegates to viper.GetBool
func (v *ViperProvider) Bool(name string) bool {
	return v.Viper.GetBool(name)
}

// StringSlice delegates to viper.GetStringSlice
func (v *ViperProvider) StringSlice(name string) []string {
	return v.Viper.GetStringSlice(name)
}
