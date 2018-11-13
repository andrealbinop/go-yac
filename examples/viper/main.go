package main

import (
	"flag"
	"fmt"
	"github.com/andrealbinop/go-yac/pkg/config"
	"log"
	"strings"
)

func main() {
	flag.Usage()
	name := flag.String("n", config.EmptyString, "Configuration Name")
	dir := flag.String("f", config.EmptyString, "Configuration Directory")
	flag.Parse()
	if *name == config.EmptyString {
		flag.Usage()
		return
	}
	loader := ViperLoader{
		Name: *name,
		Path: *dir,
	}
	cfg, err := loader.Load()
	if err != nil {
		log.Panic(err)
	}
	var entries []string
	for k, v := range cfg.AllSettings() {
		entries = append(entries, fmt.Sprintf("%v: %v", k, v))
	}
	log.Printf("Config [%v]:\n\t- %v\n", cfg.Source(), strings.Join(entries, "\n\t- "))
}
