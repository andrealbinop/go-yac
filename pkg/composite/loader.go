// Package composite allows composition of config.Loader
package composite

import (
	"errors"
	"fmt"
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/andrealbinop/go-yac/pkg/logger"
)

// Loader is a config.Loader implementation to load multiple config.Provider and merge
// them with composite.Loader
type Loader struct {
	Entries []LoaderEntry
}

// LoaderEntry holds the config.Loader included in chain loading and an Optional flag to
// allow loader to be ignored if fails to load
type LoaderEntry struct {
	// Entry is the config.Loader to help composing the config.Provider
	Entry config.Loader
	// Optional flag allowing failures to be ignored
	Optional bool
}

// Load implementation iterates through Loader.Entries to create an array of config.Provider
// results to then return a merged config.Provider.
func (b *Loader) Load() (config.Provider, error) {
	loaderCount := len(b.Entries)
	if loaderCount == config.Zero {
		return nil, errors.New("no loaders provided")
	}
	loadedCount := config.Zero
	var providers []config.Provider
	for i, ds := range b.Entries {
		loadedCount = i + 1
		progress := fmt.Sprintf("%v/%v", loadedCount, loaderCount)
		source := ds.Entry.Source()
		logger := logger.New(source)
		if cfg, err := ds.Entry.Load(); err != nil {
			if !ds.Optional {
				return nil, err
			}
			loadedCount = loadedCount - 1
			logger.Printf("failed to load (%v). Cause: %v\n", progress, err.Error())
		} else {
			logger.Printf("loaded (%v)\n", progress)
			providers = append(providers, cfg)
		}
	}
	if loadedCount == config.Zero {
		return nil, errors.New("all loaders failed")
	}
	return &Provider{Sources: providers}, nil
}
