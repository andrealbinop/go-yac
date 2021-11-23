package yaml

import (
	"github.com/andrealbinop/go-yac/internal/samples"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"testing/iotest"
)

func TestParseReadFailure(t *testing.T) {
	reader := iotest.TimeoutReader(samples.Asset("config.yml"))
	parser := Parser{}
	cfg, err := parser.Parse(reader)
	assert.Nil(t, cfg)
	assert.EqualError(t, err, "timeout")
}

func TestParseFailure(t *testing.T) {
	reader := strings.NewReader("invalid")
	parser := Parser{}
	cfg, err := parser.Parse(reader)
	assert.Nil(t, cfg)
	assert.EqualError(t, err, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `invalid` into map[string]interface {}")
}

func TestParseWithNestedProperties(t *testing.T) {
	reader := samples.Asset("config.yml")
	parser := Parser{}
	data, err := parser.Parse(reader)
	if assert.NotNil(t, data) {
		assert.Nil(t, err)
		assert.Equal(t, map[string]interface{}{
			"data.string":             "string",
			"data.int":                1,
			"data.float":              1.0,
			"data.bool":               true,
			"data.stringSlice":        []interface{}{"string1", "string2"},
			"data.intSlice":           []interface{}{400, 423},
			"data.nested.string":      "nested.string",
			"data.nested.int":         2,
			"data.nested.float":       2.0,
			"data.nested.bool":        true,
			"data.nested.stringSlice": []interface{}{"nested.string1", "nested.string2"},
			"data.nested.intSlice":    []interface{}{500, 523},
		}, data)
	}
}
