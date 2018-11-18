# go-yac [![GoDoc](https://godoc.org/github.com/andrealbinop/go-yac?status.svg)](https://godoc.org/github.com/andrealbinop/go-yac) [![Build Status](https://travis-ci.org/andrealbinop/go-yac.svg?branch=master)](https://travis-ci.org/andrealbinop/go-yac) [![Coverage Status](https://coveralls.io/repos/github/andrealbinop/go-yac/badge.svg?branch=master)](https://coveralls.io/github/andrealbinop/go-yac?branch=master) [![Go Report](https://goreportcard.com/badge/github.com/golang/dep)](https://goreportcard.com/report/github.com/andrealbinop/go-yac)

**Y**et **A**nother **C**onfig provisioning library is strongly inspired by other similar projects, such as [olebedev/config][], [micro/go-config][] and [spf13/viper][]. Why not use those aforementioned libraries you may ask? You can still use, and that's the point. Sometimes projects start small and a simple YAML configuration file is enough. However, when they grow and you need more features, you may have problems adapting your code to use another library. Through a concise set of interfaces, the main goal of this project is to minimize the impact in your application's code when you need to change configuration provisioning code. Check [here][design] for more detail.

## Installation

* With `go get`:

```bash
go get -u github.com/andrealbinop/go-yac
```

* With [dep][go-dep]:

```bash
dep ensure -add github.com/andrealbinop/go-yac
```

## Usage

* Load your configuration provider (check [here for examples][examples]) and retrieve values from it:

* Retrieve a value from the [configuration provider][]:

```go
loader := ... // build your own config.Loader instance
var err error
var cfg config.Provider
if cfg, err = loader.Load(); err != nil {
    // error handling
}
// retrieve a value associated with "key"
value := cfg.String("key")
```

## Development Notes

- This project follows [golang-standards/project-layout][] project structure.
- It's still in a proof of concept phase, calibrate your expectations.
- Please read [mocking section][] regarding testing with `config` package [interfaces][configuration provider].

[micro/go-config]: https://github.com/micro/go-config
[olebedev/config]: https://github.com/olebedev/config
[spf13/viper]: https://github.com/spf13/viper
[go-dep]: https://github.com/golang/dep
[design]: docs/DESIGN.md
[examples]: examples/README.md
[mocking section]: docs/MOCKS.md
[configuration provider]: pkg/config/config.go
[golang-standards/project-layout]: https://github.com/golang-standards/project-layout
