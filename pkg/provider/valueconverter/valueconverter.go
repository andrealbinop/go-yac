package valueconverter

import (
	"fmt"
	"strconv"
)

// Default implements config.ValueConverter
type Default struct {}

const (
	base10    = 10
	bitSize64 = 64
)

// ToString parse a value into string. It's used the package fmt to parse the value into string
func (p *Default) ToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// ToInt parses a value into int. It's used the function strconv.ParseInt to parse the value into int
func (p *Default) ToInt(value interface{}) int {
	result, ok := value.(int)
	if ok {
		return result
	}
	var resultInt64 int64
	rawValue, ok := value.(string)
	if ok {
		resultInt64, _ = strconv.ParseInt(rawValue, base10, bitSize64)
	}
	return int(resultInt64)
}

// ToFloat parses a value into float64. It's used the function strconv.ParseFloat to parse the value into float
func (p *Default) ToFloat(value interface{}) float64 {
	result, ok := value.(float64)
	if ok {
		return result
	}
	rawValue, ok := value.(string)
	if ok {
		result, _ = strconv.ParseFloat(rawValue, bitSize64)
	}
	return result
}

// ToBool parses a value into bool. It's used the function strconv.ParseBool to parse the value into bool
func (p *Default) ToBool(value interface{}) bool {
	result, ok := value.(bool)
	if ok {
		return result
	}
	rawValue, ok := value.(string)
	if ok {
		result, _ = strconv.ParseBool(rawValue)
	}
	return result
}

// ToStringSlice parses a value into string slice
func (p *Default) ToStringSlice(value interface{}) []string {
	var result []string
	result, ok := value.([]string)
	if ok {
		return result
	}
	rawValue, ok := value.([]interface{})
	if ok {
		for _, str := range rawValue {
			result = append(result, p.ToString(str))
		}
	}
	return result
}

// ToIntSlice parses a value into int slice
func (p *Default) ToIntSlice(value interface{}) []int {
	var result []int
	result, ok := value.([]int)
	if ok {
		return result
	}
	rawValue, ok := value.([]interface{})
	if ok {
		for _, intValue := range rawValue {
			result = append(result, p.ToInt(intValue))
		}
	}
	return result
}
