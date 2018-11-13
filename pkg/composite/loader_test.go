package composite

import (
	"errors"
	"github.com/andrealbinop/go-yac/internal/mocks"
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/andrealbinop/go-yac/pkg/logger"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadLoaderWithoutLoaders(t *testing.T) {
	Loader := &Loader{}
	cfg, err := Loader.Load()
	assert.Nil(t, cfg)
	assert.EqualError(t, err, "no loaders provided")
}

func TestLoadLoaderWithFailedRequiredLoader(t *testing.T) {
	expectedError := errors.New("simulate loader entry fail")
	loader := mockLoader("mock source", nil, expectedError)
	Loader := &Loader{
		Entries: []LoaderEntry{
			{
				Entry:    loader,
				Optional: false,
			},
		},
	}
	cfg, err := Loader.Load()
	if loader.AssertExpectations(t) {
		assert.Nil(t, cfg)
		assert.Equal(t, expectedError, err)
	}
}

func TestLoadLoaderWithAllLoadersFailed(t *testing.T) {
	expectedError := errors.New("simulate loader entry fail")
	loader := mockLoader("mock source", nil, expectedError)
	Loader := &Loader{
		Entries: []LoaderEntry{
			{
				Entry:    loader,
				Optional: true,
			},
		},
	}
	logger := mockLogger()
	logger.On("Printf", "failed to load (%v). Cause: %v\n", "1/1", expectedError.Error())
	cfg, err := Loader.Load()
	if loader.AssertExpectations(t) {
		assert.Nil(t, cfg)
		assert.EqualError(t, err, "all loaders failed")
	}
}

func TestLoadLoaderWithMultipleLoadersOk(t *testing.T) {
	providerA := &mocks.Provider{}
	providerA.On("Source").Return("mock source A")
	providerB := &mocks.Provider{}
	providerB.On("Source").Return("mock source B")
	loaderA := mockLoader("mock source A", providerA, nil)
	loaderB := mockLoader("mock source B", providerB, nil)
	loader := &Loader{
		Entries: []LoaderEntry{
			{Entry: loaderA},
			{Entry: loaderB},
		},
	}
	logger := mockLogger()
	logger.On("Printf", "loaded (%v)\n", "1/2")
	logger.On("Printf", "loaded (%v)\n", "2/2")
	cfg, err := loader.Load()
	loaderA.AssertExpectations(t)
	loaderB.AssertExpectations(t)
	assert.NoError(t, err)
	if assert.NotNil(t, cfg) {
		assert.Equal(t, "mock source A>mock source B", cfg.Source())
		providerA.AssertExpectations(t)
		providerB.AssertExpectations(t)
	}
}

func mockLoader(name string, cfg config.Provider, err error) *mocks.Loader {
	loader := &mocks.Loader{}
	loader.On("Source").Return(name)
	loader.On("Load").Return(cfg, err)
	return loader
}

func mockLogger() *mocks.Logger {
	newMockLogger := &mocks.Logger{}
	logger.Factory = func(string) logger.Logger {
		return newMockLogger
	}
	return newMockLogger
}
