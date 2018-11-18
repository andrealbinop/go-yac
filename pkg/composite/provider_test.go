package composite

import (
	"github.com/andrealbinop/go-yac/internal/mocks"
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	nonExistent = "non_existent"
	existent    = "existent"
	valString   = "string"
	valInt      = 1
	valFloat    = float64(valInt)
	valBool     = true
	mockSourceA = "mockSourceA"
	mockSourceB = "mockSourceB"
)

func TestPropertyNotFoundString(t *testing.T) {
	provider := Provider{}
	assert.Empty(t, provider.String(nonExistent))
}

func TestPropertyNotFoundInt(t *testing.T) {
	provider := Provider{}
	assert.Zero(t, provider.Int(nonExistent))
}

func TestPropertyNotFoundFloat(t *testing.T) {
	provider := Provider{}
	assert.Zero(t, provider.Float(nonExistent))
}

func TestPropertyNotFoundBool(t *testing.T) {
	provider := Provider{}
	assert.False(t, provider.Bool(nonExistent))
}

func TestPropertyNotFoundStringSlice(t *testing.T) {
	provider := Provider{}
	assert.Empty(t, provider.StringSlice(nonExistent))
}

func TestPropertyNotFoundGet(t *testing.T) {
	provider := Provider{}
	result, ok := provider.Get(nonExistent)
	assert.Nil(t, result)
	assert.False(t, ok)
}

func TestPropertyNotSet(t *testing.T) {
	provider := Provider{}
	assert.False(t, provider.IsSet(nonExistent))
}

func TestNoProperties(t *testing.T) {
	provider := Provider{}
	assert.Empty(t, provider.AllSettings())
}

func TestNoSources(t *testing.T) {
	provider := Provider{}
	assert.Empty(t, provider.Source())
}

func TestSingleSourceName(t *testing.T) {
	childProvider := &mocks.Provider{}
	childProvider.On("Source").Return(mockSourceA)
	provider := Provider{
		Sources: []config.Provider{childProvider},
	}
	assert.Equal(t, mockSourceA, provider.Source())
}

func TestMultipleSourcesName(t *testing.T) {
	parentProvider := &mocks.Provider{}
	parentProvider.On("Source").Return(mockSourceA)
	childProvider := &mocks.Provider{}
	childProvider.On("Source").Return(mockSourceB)
	provider := Provider{
		Sources: []config.Provider{parentProvider, childProvider},
	}
	assert.Equal(t, mockSourceA+">"+mockSourceB, provider.Source())
}

func TestAllSettingsMultipleSources(t *testing.T) {
	parentProvider := &mocks.Provider{}
	parentProvider.On("AllSettings").Return(map[string]interface{}{
		"prop.1": "parent.1",
		"prop.2": "parent.2",
	})
	childProvider := &mocks.Provider{}
	childProvider.On("AllSettings").Return(map[string]interface{}{
		"prop.2": "child.2",
		"prop.3": "child.3",
	})
	expected := map[string]interface{}{
		"prop.1": "parent.1",
		"prop.2": "child.2",
		"prop.3": "child.3",
	}
	provider := Provider{
		Sources: []config.Provider{parentProvider, childProvider},
	}
	assert.Equal(t, expected, provider.AllSettings())

}

func TestSetAndGetPropertyMultipleSources(t *testing.T) {
	parentProvider := mockProviderWithCall("Get", valInt, true)
	childProvider := mockProviderWithCall("Get", nil, false)
	childProvider.On("Set", existent, valString).Return(nil)
	provider := Provider{
		Sources: []config.Provider{parentProvider, childProvider},
	}
	assert.Equal(t, valInt, provider.Set(existent, valString))
}

func TestStringPropertyInParentProvider(t *testing.T) {
	parentProvider := mockProviderWithCall("String", valString)
	childProvider := mockProviderWithCall("String", nil)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, valString, provider.String(existent))
	})
}

func TestIntPropertyInParentProvider(t *testing.T) {
	parentProvider := mockProviderWithCall("Int", valInt)
	childProvider := mockProviderWithCall("Int", nil)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, valInt, provider.Int(existent))
	})
}

func TestFloatPropertyInParentProvider(t *testing.T) {
	parentProvider := mockProviderWithCall("Float", valFloat)
	childProvider := mockProviderWithCall("Float", nil)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, valFloat, provider.Float(existent))
	})
}

func TestBoolPropertyInParentProvider(t *testing.T) {
	parentProvider := mockProviderWithCall("Bool", valBool)
	childProvider := mockProviderWithCall("Bool", nil)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, valBool, provider.Bool(existent))
	})
}

func TestStringSlicePropertyInParentProvider(t *testing.T) {
	value := []string{valString}
	parentProvider := mockProviderWithCall("StringSlice", value)
	childProvider := mockProviderWithCall("StringSlice", nil)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, value, provider.StringSlice(existent))
	})
}

func TestStringPropertyInChildProvider(t *testing.T) {
	parentProvider := &mocks.Provider{}
	childProvider := mockProviderWithCall("String", valString)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, valString, provider.String(existent))
	})
}

func TestIntPropertyInChildProvider(t *testing.T) {
	parentProvider := &mocks.Provider{}
	childProvider := mockProviderWithCall("Int", valInt)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, valInt, provider.Int(existent))
	})
}

func TestFloatPropertyInChildProvider(t *testing.T) {
	parentProvider := &mocks.Provider{}
	childProvider := mockProviderWithCall("Float", valFloat)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, valFloat, provider.Float(existent))
	})
}

func TestBoolPropertyInChildProvider(t *testing.T) {
	parentProvider := &mocks.Provider{}
	childProvider := mockProviderWithCall("Bool", valBool)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, valBool, provider.Bool(existent))
	})
}

func TestStringSlicePropertyInChildProvider(t *testing.T) {
	value := []string{valString}
	parentProvider := &mocks.Provider{}
	childProvider := mockProviderWithCall("StringSlice", value)
	assertCompositeProvider(t, parentProvider, childProvider, func(provider config.Provider) {
		assert.Equal(t, value, provider.StringSlice(existent))
	})
}

func mockProviderWithCall(methodName string, value interface{}, additionalReturn ...interface{}) *mocks.Provider {
	provider := mocks.Provider{}
	hasValue := value != nil
	provider.On("IsSet", existent).Return(hasValue)
	if hasValue {
		returnArgs := append([]interface{}{value}, additionalReturn...)
		provider.On(methodName, existent).Return(returnArgs...)
	}
	return &provider

}

func assertCompositeProvider(t *testing.T, parent *mocks.Provider, child *mocks.Provider, assertion func(config.Provider)) {
	provider := Provider{
		Sources: []config.Provider{parent, child},
	}
	assertion(&provider)
	parent.AssertExpectations(t)
	child.AssertExpectations(t)
}
