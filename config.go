package main

import (
	"flag"
	"net/url"
	"os"
	"strconv"
)

type Config struct {
	BaseURL string
	Silent  bool
	Debug   bool
}

func loadEnv(config *Config) {
	if u, e := url.Parse(os.Getenv("LOGPASTA_URL")); e == nil {
		config.BaseURL = u.String()
	}
	if v, e := strconv.ParseBool(os.Getenv("LOGPASTA_SILENT")); e == nil {
		config.Silent = v
	}
	if v, e := strconv.ParseBool(os.Getenv("LOGPASTA_DEBUG")); e == nil {
		config.Debug = v
	}
}

func loadFlags(config *Config) {
	flag.StringVar(&config.BaseURL, "u", config.BaseURL, "logpasta base url, without trailing slash")
	flag.BoolVar(&config.Silent, "s", config.Silent, "silent mode - suppress logs unless request fails")
	flag.BoolVar(&config.Debug, "d", config.Debug, "debug mode - output debug information")
	flag.Parse()
}
