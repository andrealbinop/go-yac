package main

import (
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/spf13/viper"
)

type ViperLoader struct {
	Name string
	Path string
}

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
