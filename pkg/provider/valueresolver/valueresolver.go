package valueresolver

import "github.com/andrealbinop/go-yac/pkg/config"

// Default implements config.ValueResolver
type Default struct {
}

// IsSet returns if there's a value associated with the name of repository.
func (k *Default) IsSet(name string, repository config.Repository) bool {
	return repository.IsSet(name)
}

// Resolve returns the value associated with the name of repository.
func (k *Default) Resolve(name string, repository config.Repository) (interface{}, bool) {
	return repository.Get(name)
}
