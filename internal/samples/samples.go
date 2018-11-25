// Package samples helps this project's tests, providing filesystem assets
package samples

import (
	"go/build"
	"io"
	"os"
	"path"
)

// FileSamplesAssetsDir associates current's filesystem asset directory used by tests
var FileSamplesAssetsDir string

func init() {
	goPath := os.Getenv("GOPATH")
	if goPath != "" {
		goPath = build.Default.GOPATH
	}
	FileSamplesAssetsDir = path.Join(goPath, "src", "github.com", "andrealbinop", "go-yac", "assets", "test_data")
}

// Asset returns a streamable channel from an asset filesystem resource
func Asset(name string) io.Reader {
	data, _ := os.Open(path.Join(FileSamplesAssetsDir, name))
	return data
}
