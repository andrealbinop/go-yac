# go-yac/mocks

This projects makes use of [vektra/mockery][] to mock [config][] interfaces for unit testing. If any interface changes occurs, you'll need to update the mocks with the following command ([mockery cli][vektra/mockery#installation] will need to be installed):

```bash
mockery -all -case underscore -dir ./pkg -output ./test/mocks
```

[vektra/mockery]: https://github.com/vektra/mockery
[vektra/mockery#installation]: https://github.com/vektra/mockery#installation
[config]: pkg/config/config.go