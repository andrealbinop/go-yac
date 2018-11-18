# go-yac/design

This section tries to shed some light regarding some of the design decisions for this project.

The main go is to provide concise `configuration provisioning` interfaces, so when features are added or dependencies changed you don't need to change code that consumes configuration data or trigger them to be loaded.

First we need to estabilish, what exactly is `configuration provisioning`? This term is too broad, maybe there's a better name for it, but in the scope of this project means the following:

```
Retrieve or provide structured data to allow application's to work in different contexts. Credentials, feature enable flags are, though not limited to, examples of configuration data.
```

As a configuration consumer, you should:

- Set the sources from where configuration should be loaded.
- Trigger configuration to be loaded (or reloaded).
- Consume the configuration.

For each interaction, a set of interfaces and default implementations are provided, and the following sections details each group.

## Loading configuration

The interface `config.Loader` encompasses configuration load phase, and can be used to create adapters for totally different `configuration provisioning`
solutions, as [exemplified here][viper]. However, a few additional helper interfaces to highlight common loading phases, such as reading from a source, parsing data, compose and merge them into a provider. Below a [class diagram][] for all loading related interfaces and utilities:

![loader.svg][]

[viper]: ../examples/viper/README.md
[class diagram]: https://en.wikipedia.org/wiki/Class_diagram
[loader.svg]: ../assets/uml/loader.svg

## Consuming configuration

TODO...