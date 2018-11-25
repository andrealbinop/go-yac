// Package loader provides reading configuration from os environment
package loader

import (
	"fmt"
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/andrealbinop/go-yac/pkg/provider"
	"github.com/andrealbinop/go-yac/pkg/repository"
	"os"
	"strings"
)

const (
	// EnvSource constant for environment variable source
	EnvSource = "env"
	// ArgsSource constant for command line args source
	ArgsSource = "args"
)

// EqualDelimiterParser parses Data string slice to a config.Provider, considering only prefixed variables
type EqualDelimiterParser struct {
	// Name identifier for the data that is being loaded
	Name string
	// Prefix to filter undesirable entries in slice
	Prefix string
	// Data to be parsed
	Data []string
}

// Load config.Provider instance from the provided Data string slice and prefix
func (p *EqualDelimiterParser) Load() (cfg config.Provider, err error) {
	envData := make(map[string]interface{})
	for _, e := range p.Data {
		if !strings.HasPrefix(e, p.Prefix) {
			continue
		}
		pair := strings.Split(e, "=")
		key := strings.TrimPrefix(pair[0], p.Prefix)
		envData[key] = pair[1]
	}
	cfg = &provider.Default{
		Repository: &repository.Map{
			Database: envData,
		},
		SourceName: p.Source(),
	}
	return
}

// Source returns an identifier for this loader
func (p *EqualDelimiterParser) Source() string {
	return fmt.Sprintf("%v:%v*", p.Name, p.Prefix)
}

// Env creates a loader for environment variables with prefixed keys
func Env(prefix string) config.Loader {
	return &EqualDelimiterParser{
		Name:   EnvSource,
		Prefix: prefix,
		Data:   os.Environ(),
	}
}

// Args creates a loader for environment variables with prefixed keys
func Args(prefix string) config.Loader {
	return &EqualDelimiterParser{
		Name:   ArgsSource,
		Prefix: prefix,
		Data:   os.Args,
	}
}
