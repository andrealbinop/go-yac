package reader

import (
	"github.com/andrealbinop/go-yac/pkg/config"
	"io"
	"os"
	"path"
)

// File implements config.Loader to provide a streamable channel from filesystem sources
type File struct {
	// ParentDir to lookup for a filesystem resource
	ParentDir string
}

// Read returns a io.Reader to the provided filesystem location
func (r *File) Read(location string) (io.Reader, error) {
	if r.ParentDir != config.EmptyString {
		location = path.Join(r.ParentDir, location)
	}
	return os.Open(location)
}
