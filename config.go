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
	Timeout int64
}

func loadEnv(config *Config) {
	if u, e := url.Parse(os.Getenv("LOGPASTA_URL")); e == nil && len(u.String()) > 0 {
		config.BaseURL = u.String()
	}
	if v, e := strconv.ParseBool(os.Getenv("LOGPASTA_SILENT")); e == nil {
		config.Silent = v
	}
	if v, e := strconv.ParseBool(os.Getenv("LOGPASTA_DEBUG")); e == nil {
		config.Debug = v
	}
	if v, e := strconv.ParseInt(os.Getenv("LOGPASTA_TIMEOUT"), 10, 64); e == nil {
		config.Timeout = v
	}
}

func loadFlags(config *Config) {
	flag.StringVar(&config.BaseURL, "u", config.BaseURL, "logpasta base url, without trailing slash")
	flag.BoolVar(&config.Silent, "s", config.Silent, "silent mode - suppress logs unless request fails")
	flag.BoolVar(&config.Debug, "d", config.Debug, "debug mode - output debug information")
	flag.Int64Var(&config.Timeout, "t", config.Timeout, "timeout - http client timeout in seconds")
	flag.Parse()
}
