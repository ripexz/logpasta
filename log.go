package main

import (
	"flag"
	"fmt"
	"log"
)

type logger struct{}

func (writer logger) Write(bytes []byte) (int, error) {
	return fmt.Print(string(bytes))
}

func initLogger() {
	log.SetFlags(0)
	log.SetOutput(new(logger))
}

func printVersion() {
	log.Printf("Logpasta CLI %s", version)
}

func printUsage() {
	var flags string
	flag.VisitAll(func(f *flag.Flag) {
		flags += f.Name
	})

	log.Printf(
		"Usage of logpasta:\n  logpasta [-%s] [text ...]\n  [command] | logpasta [-%s]\n  logpasta [-%s] < [file]\n\n",
		flags, flags, flags,
	)

	log.Println("Options:")
	flag.PrintDefaults()
}

func printDebugInfo(conf *Config) {
	log.Printf("Running with config:\n - BaseURL: %s\n - Silent: %v\n - Debug: %v\n - Timeout: %d\n - Copy: %v",
		conf.BaseURL,
		conf.Silent,
		conf.Debug,
		conf.Timeout,
		conf.Copy,
	)
}
