# go-yac/examples/viper

This sample show how to compose a configuration provider from [spf13/viper] library. This example seems silly, putting interfaces on top a well known library. However, if you start small (with local YAML configuration) and needs to migrate or add more complex `configuration provisioning` features. This example show how could you do that without changing much code that triggers configuration loading and consumes it's data.

## Usage

```bash
go run main.go provider.go loader.go -n ${CONFIG_NAME} -f ${CONFIG_DIR}
```

## Example

```bash
go run main.go provider.go loader.go -n config -f ../../assets/test_data
Usage of /tmp/___viper:
2018/11/18 16:35:41 Config [/home/andre/Code/Go/src/github.com/andrealbinop/go-yac/test/samples/assets/config.yml]:
	- data: map[int:1 float:1 string:string stringslice:[string1 string2] bool:true nested:map[string:nested.string int:2 float:2 bool:true stringslice:[nested.string1 nested.string2]]]
```

[spf13/viper]: https://github.com/spf13/viper
