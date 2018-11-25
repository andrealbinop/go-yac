// Package go_yac is a configuration provisioning toolkit, inspired by other libraries such as viper, go-config, among others.
//
// go_yac contains the following packages:
//
// The composite package allows composing hierarchically config.Loader and config.Provider implementations.
//
// The config package provides a cohesive set of interfaces to expose configuration provisioning features.
//
// The loader package contains useful implementations for config.Loader interface, relying only on go stdlib.
//
// The logger package provides a stdlib log compatible interface and mechanisms to easily swap implementations.
//
// The provider package offers a default config.Repository backed implementation for config.Provider.
//
// The reader package contains useful implementations for loader.Reader interface, relying only on go stdlib.
//
// The yaml package provides a go-yaml based loader.Parser implementation
package go_yac

// blank imports help docs.
import (
	// composite package
	_ "github.com/andrealbinop/go-yac/pkg/composite"
	// config package
	_ "github.com/andrealbinop/go-yac/pkg/config"
	// loader package
	_ "github.com/andrealbinop/go-yac/pkg/loader"
	// logger package
	_ "github.com/andrealbinop/go-yac/pkg/logger"
	// provider package
	_ "github.com/andrealbinop/go-yac/pkg/provider"
	// reader package
	_ "github.com/andrealbinop/go-yac/pkg/reader"
	// yaml package
	_ "github.com/andrealbinop/go-yac/pkg/parser/yaml"
)
