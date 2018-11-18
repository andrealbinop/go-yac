# go-yac/examples/yaml

This sample show how to compose a configuration provider from YAML files. Since multiple YAML files can be
informed, they will be merged, and the last file will have higher precedence from the previous.

## Usage

```bash
go run main.go -f ${COMMA_DELIMITED_YAML_FILES}
```

## Example

```bash
$ go run main.go -f ../../assets/test_data

yac.data:test/samples/assets/config.yml loaded (1/1)
Config [data:test/samples/assets/config.yml]:
	- data.float: 1
	- data.nested.int: 2
	- data.string: string
	- data.nested.float: 2
	- data.int: 1
	- data.bool: true
	- data.stringSlice: [string1 string2]
	- data.nested.stringSlice: [nested.string1 nested.string2]
	- data.nested.string: nested.string
	- data.nested.bool: true
```

