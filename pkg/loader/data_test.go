package loader

import (
	"errors"
	"github.com/andrealbinop/go-yac/internal/mocks"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

const (
	valLocation = "mock location"
	valData     = "mock data"
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
	mockReader := strings.NewReader(valData)
	reader.On("Read", valLocation).Return(mockReader, nil)
	parser := mocks.Parser{}
	expectedError := errors.New("mock read error")
	parser.On("Parse", mockReader).Return(nil, expectedError)
	loader := Data{
		Location: valLocation,
		Reader:   &reader,
		Parser:   &parser,
	}
	cfg, err := loader.Load()
	reader.AssertExpectations(t)
	assert.Nil(t, cfg)
	assert.Equal(t, expectedError, err)
}

func TestDataLoad(t *testing.T) {
	reader := mocks.Reader{}
	mockReader := strings.NewReader(valData)
	reader.On("Read", valLocation).Return(mockReader, nil)
	parser := mocks.Parser{}
	database := map[string]interface{}{
		"prop.1": "value.1",
	}
	parser.On("Parse", mockReader).Return(database, nil)
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
