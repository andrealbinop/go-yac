package main

import "github.com/spf13/viper"

type ViperProvider struct {
	ConfigName string
	Viper      *viper.Viper
}

func (v *ViperProvider) Source() string {
	return v.Viper.ConfigFileUsed()
}

func (v *ViperProvider) Set(name string, value interface{}) interface{} {
	previous, _ := v.Get(name)
	viper.Set(name, value)
	return previous
}

func (v *ViperProvider) Get(name string) (interface{}, bool) {
	return v.Viper.Get(name), v.IsSet(name)
}

func (v *ViperProvider) IsSet(name string) bool {
	return v.Viper.IsSet(name)
}

func (v *ViperProvider) AllSettings() map[string]interface{} {
	return v.Viper.AllSettings()
}

func (v *ViperProvider) String(name string) string {
	return v.Viper.GetString(name)
}

func (v *ViperProvider) Int(name string) int {
	return v.Viper.GetInt(name)
}

func (v *ViperProvider) Float(name string) float64 {
	return v.Viper.GetFloat64(name)
}

func (v *ViperProvider) Bool(name string) bool {
	return v.Viper.GetBool(name)
}

func (v *ViperProvider) StringSlice(name string) []string {
	return v.Viper.GetStringSlice(name)
}
