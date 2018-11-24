package loader

import (
	"errors"
	"github.com/andrealbinop/go-yac/internal/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

const (
	valLocation = "mock location"
)

func TestDataFailedToRead(t *testing.T) {
	reader := mocks.Reader{}
	expectedError := errors.New("mock read error")
	reader.On("Read", valLocation).Return(nil, expectedError)
	loader := Data{
		Location: valLocation,
		Reader:   &reader,
	}
	cfg, err := loader.Load()
	reader.AssertExpectations(t)
	assert.Nil(t, cfg)
	assert.Equal(t, expectedError, err)
}

func TestDataFailedToParse(t *testing.T) {
	reader := mocks.Reader{}
	ioReader := &mocks.IOReadCloser{}
	reader.On("Read", valLocation).Return(ioReader, nil)
	ioReader.On("Close").Return(nil)
	parser := mocks.Parser{}
	expectedError := errors.New("mock read error")
	parser.On("Parse", ioReader).Return(nil, expectedError)
	loader := Data{
		Location: valLocation,
		Reader:   &reader,
		Parser:   &parser,
	}
	cfg, err := loader.Load()
	reader.AssertExpectations(t)
	ioReader.AssertExpectations(t)
	assert.Nil(t, cfg)
	assert.Equal(t, expectedError, err)
}

func TestDataFailedToClose(t *testing.T) {
	reader := mocks.Reader{}
	ioReader := &mocks.IOReadCloser{}
	expectedError := errors.New("mock close error")
	reader.On("Read", valLocation).Return(ioReader, nil)
	ioReader.On("Close").Return(expectedError)
	parser := mocks.Parser{}
	parser.On("Parse", ioReader).Return(make(map[string]interface{}), nil)
	loader := Data{
		Location: valLocation,
		Reader:   &reader,
		Parser:   &parser,
	}
	_, err := loader.Load()
	reader.AssertExpectations(t)
	ioReader.AssertExpectations(t)
	assert.Equal(t, expectedError, err)
}

func TestDataLoadForIOReader(t *testing.T) {
	ioReader := &mocks.IOReader{}
	testDataLoad(t, ioReader)
	ioReader.AssertExpectations(t)
}

func TestDataLoadForIOReadCloser(t *testing.T) {
	ioReader := &mocks.IOReadCloser{}
	ioReader.On("Close").Return(nil)
	testDataLoad(t, ioReader)
	ioReader.AssertExpectations(t)
}

func testDataLoad(t *testing.T, ioReaderResponse io.Reader) {
	reader := mocks.Reader{}
	reader.On("Read", valLocation).Return(ioReaderResponse, nil)
	parser := mocks.Parser{}
	database := map[string]interface{}{
		"prop.1": "value.1",
	}
	parser.On("Parse", ioReaderResponse).Return(database, nil)
	loader := Data{
		Location: valLocation,
		Reader:   &reader,
		Parser:   &parser,
	}
	cfg, err := loader.Load()
	reader.AssertExpectations(t)
	assert.NoError(t, err)
	assert.Equal(t, database, cfg.AllSettings())
	assert.Equal(t, "data:"+valLocation, cfg.Source())
}
