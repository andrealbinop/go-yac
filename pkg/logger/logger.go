// Package logger provides an interfaces used to log configuration provisioning default implementations.
// This interface is compatible with stdlib log.Logger (which doesn't have an interface).
// You may provide your own logger.Logger implementation by swapping logger.Factory provider
package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
)

const (
	// DefaultPrefix constant for all triggered logs
	DefaultPrefix = "yac"
)

// A Logger interface compatible with log.Logger logging instructions, which currently don't provide an interface
type Logger interface {
	// Print calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Print.
	Print(...interface{})
	// Printf calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Printf.
	Printf(string, ...interface{})
	// Println calls l.Output to print to the logger.
	// Arguments are handled in the manner of fmt.Println.
	Println(...interface{})
	// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
	Fatal(...interface{})
	// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
	Fatalf(string, ...interface{})
	// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
	Fatalln(...interface{})
	// Panic is equivalent to l.Print() followed by a call to panic().
	Panic(...interface{})
	// Panicf is equivalent to l.Printf() followed by a call to panic().
	Panicf(string, ...interface{})
	// Panicln is equivalent to l.Println() followed by a call to panic().
	Panicln(...interface{})
}

// Factory function creates a Logger implementation for a given prefix. May be swapped for your own. By default
// uses log.New(os.Stderr, prefix, log.LstdFlags)
var Factory func(string) Logger

func init() {
	Factory = func(prefix string) Logger {
		return log.New(os.Stderr, prefix, log.LstdFlags)
	}
}

// New function returns a Logger implementation for a given prefix
func New(prefix string) Logger {
	if prefix != "" {
		prefix = "." + prefix
	}
	prefix = fmt.Sprintf("%v%v ", DefaultPrefix, strings.TrimSpace(prefix))
	return Factory(prefix)
}
