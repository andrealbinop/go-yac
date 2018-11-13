// yaml package parses yaml data to be used by map configuration databases
package yaml

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"io/ioutil"
)

// Parser implements loader.Parser interface
type Parser struct {
}

// Parse function reads io.Reader into a map[string]interface{} with yaml parsed information
func (p *Parser) Parse(reader io.Reader) (result map[string]interface{}, err error) {
	var data []byte
	if data, err = ioutil.ReadAll(reader); err != nil {
		return
	}
	parsed := make(map[string]interface{})
	if err = yaml.Unmarshal(data, &parsed); err != nil {
		return
	}
	result = make(map[string]interface{})
	for k, v := range parsed {
		flattenProperties(k, v, result)
	}
	return
}

func flattenProperties(name string, value interface{}, result map[string]interface{}) {
	if mapValue, ok := value.(map[interface{}]interface{}); ok {
		for k, v := range mapValue {
			fullKey := fmt.Sprintf("%v.%v", name, k)
			flattenProperties(fullKey, v, result)
		}
	} else {
		result[name] = value
	}
}
