package main

import (
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/spf13/viper"
)

// ViperLoader is a loader implementation for config.Loader interface
type ViperLoader struct {
	// Name for viper loading configuration
	Name string
	//  Path directory to load configuration files from
	Path string
}

// Load creates a config.Provider implementation from viper source
func (l *ViperLoader) Load() (config.Provider, error) {
	viper.SetConfigName(l.Name)
	viper.AddConfigPath(l.Path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	provider := ViperProvider{
		Viper: viper.GetViper(),
	}
	return &provider, nil
}
