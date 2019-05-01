package main

import (
	"os"
	"strconv"
)

func loadEnv(config *Config) {
	if s, e := strconv.ParseBool(os.Getenv("LOGPASTA_SILENT")); e == nil {
		config.Silent = s
	}
}
