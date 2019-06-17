// Package loader provides reading configuration from streamable data
package loader

import (
	"fmt"
	"io"
	"path"

	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/andrealbinop/go-yac/pkg/provider"
	"github.com/andrealbinop/go-yac/pkg/provider/valueconverter"
	"github.com/andrealbinop/go-yac/pkg/provider/valueresolver"
	"github.com/andrealbinop/go-yac/pkg/repository"
)

// Data reads uses the the Location, Reader and Parser to load an instance of config.Provider
type Data struct {
	// Location is the address identifier from where the data will be read
	Location string
	// Reader creates a readable channel (io.Reader) from the provided location
	Reader Reader
	// Parser parses the data read from the channel to the database used by config.Provider
	Parser Parser
}

// Reader interface with function to provide the readable channel (io.ReadCloser).
type Reader interface {
	// Read provides a channel to be read from the provided location. if io.Reader response is also a io.Closer, loader.Data will try to close it after parse
	Read(string) (io.Reader, error)
}

// Parser interface with function to read data to a key value store
type Parser interface {
	// Parse provided data to to a key value store (map[string]interface{})
	Parse(io.Reader) (map[string]interface{}, error)
}

// IOReader currently just wraps io.Reader, allowing for mock generation when testing with io.Reader
type IOReader interface {
	io.Reader
}

// IOReadCloser currently just wraps io.ReadCloser, allowing for mock generation when testing with io.ReadCloser
type IOReadCloser interface {
	io.ReadCloser
}

// Load uses the Location, Reader and Parser to build config.Provider.
func (s *Data) Load() (cfg config.Provider, err error) {
	var reader io.Reader
	if reader, err = s.Reader.Read(s.Location); err != nil {
		return
	}
	var parsed map[string]interface{}
	if parsed, err = s.Parser.Parse(reader); err == nil {
		cfg = &provider.Default{
			SourceName: s.Source(),
			Repository: &repository.Map{
				Database: parsed,
			},
			ValueConverter: &valueconverter.Default{},
			ValueResolver:   &valueresolver.Default{},
		}
	}
	if closer, ok := reader.(io.Closer); ok {
		errClose := closer.Close()
		if errClose != nil {
			err = errClose
		}
	}
	return
}

// Source identifies that this is a data loader and which location is associated to
func (s *Data) Source() string {
	return fmt.Sprintf("data:%v", path.Clean(s.Location))
}
