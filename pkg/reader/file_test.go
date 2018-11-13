package reader

import (
	"github.com/andrealbinop/go-yac/internal/samples"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"testing"
)

func TestFileNotFoundReader(t *testing.T) {
	reader := File{}
	result, err := reader.Read("non_existent")
	assert.Nil(t, result)
	assert.EqualError(t, err, "open non_existent: no such file or directory")
}

func TestFileRead(t *testing.T) {
	reader := File{
		ParentDir: samples.FileSamplesAssetsDir,
	}
	result, err := reader.Read("ok.txt")
	if assert.NoError(t, err) {
		var bytes []byte
		bytes, err = ioutil.ReadAll(result)
		if assert.NoError(t, err) {
			assert.Equal(t, "ok\n", string(bytes))
		}
	}
}
