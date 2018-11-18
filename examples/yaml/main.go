package main

import (
	"flag"
	"fmt"
	"github.com/andrealbinop/go-yac/pkg/composite"
	"github.com/andrealbinop/go-yac/pkg/config"
	"github.com/andrealbinop/go-yac/pkg/loader"
	"github.com/andrealbinop/go-yac/pkg/parser/yaml"
	"github.com/andrealbinop/go-yac/pkg/reader"
	"log"
	"strings"
)

func main() {
	flag.Usage()
	files := flag.String("f", config.EmptyString, "Comma separated YAML configuration files to be loaded and merged")
	flag.Parse()
	if *files == config.EmptyString {
		flag.Usage()
		return
	}
	cfg, err := loadConfig(strings.Split(*files, ","))
	if err != nil {
		log.Panic(err)
	}

	var entries []string
	for k, v := range cfg.AllSettings() {
		entries = append(entries, fmt.Sprintf("%v: %v", k, v))
	}
	log.Printf("Config [%v]:\n\t- %v\n", cfg.Source(), strings.Join(entries, "\n\t- "))
}

func loadConfig(files []string) (config.Provider, error) {
	var loaders []composite.LoaderEntry
	for _, file := range files {
		loaders = append(loaders, composite.LoaderEntry{
			Entry: &loader.Data{
				Location: strings.TrimSpace(file),
				Reader:   &reader.File{},
				Parser:   &yaml.Parser{},
			},
		})
	}
	compositeLoader := composite.Loader{
		Entries: loaders,
	}
	return compositeLoader.Load()
}
