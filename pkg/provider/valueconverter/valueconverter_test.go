package valueconverter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString_Success(t *testing.T) {
	parser := Default{}
	expected := "10"
	assert.Equal(t, expected, parser.ToString(10))
}

func TestToInt_Success(t *testing.T) {
	parser := Default{}
	expected := 10
	assert.Equal(t, expected, parser.ToInt(10))
}

func TestToInt_WrongValue(t *testing.T) {
	parser := Default{}
	expected := 0
	assert.Equal(t, expected, parser.ToInt("a"))
}

func TestToInt_RawValue(t *testing.T) {
	parser := Default{}
	expected := 10
	var value interface{} = "10"
	assert.Equal(t, expected, parser.ToInt(value))
}

func TestToFloat_Success(t *testing.T) {
	parser := Default{}
	expected := 10.10
	assert.Equal(t, expected, parser.ToFloat(10.10))
}

func TestToFloat_WrongValue(t *testing.T) {
	parser := Default{}
	var expected float64
	assert.Equal(t, expected, parser.ToFloat("a"))
}

func TestToFloat_RawValue(t *testing.T) {
	parser := Default{}
	expected := 10.10
	var value interface{} = "10.10"
	assert.Equal(t, expected, parser.ToFloat(value))
}

func TestToBool_Success(t *testing.T) {
	parser := Default{}
	expected := true
	assert.Equal(t, expected, parser.ToBool(true))
}

func TestToBool_WrongValue(t *testing.T) {
	parser := Default{}
	expected := false
	assert.Equal(t, expected, parser.ToBool("a"))
}

func TestToBool_RawValue(t *testing.T) {
	parser := Default{}
	expected := true
	var value interface{} = "1"
	assert.Equal(t, expected, parser.ToBool(value))
}

func TestToStringSlice_Success(t *testing.T) {
	parser := Default{}
	expected := []string{"val1", "val2"}
	assert.Equal(t, expected, parser.ToStringSlice([]string{"val1", "val2"}))
}

func TestToStringSlice_WrongValue(t *testing.T) {
	parser := Default{}
	var expected []string
	assert.Equal(t, expected, parser.ToStringSlice(10))
}

func TestToStringSlice_RawValue(t *testing.T) {
	parser := Default{}
	expected := []string{"val1", "val2"}
	assert.Equal(t, expected, parser.ToStringSlice([]interface{}{"val1", "val2"}))
}

func TestToIntSlice_Success(t *testing.T) {
	parser := Default{}
	expected := []int{200, 400, 500}
	assert.Equal(t, expected, parser.ToIntSlice([]int{200, 400, 500}))
}

func TestToIntSlice_WrongValue(t *testing.T) {
	parser := Default{}
	var expected []int
	assert.Equal(t, expected, parser.ToIntSlice("val1"))
}

func TestToIntSlice_RawValue(t *testing.T) {
	parser := Default{}
	expected := []int{200, 400, 500}
	assert.Equal(t, expected, parser.ToIntSlice([]interface{}{200, 400, 500}))
}
