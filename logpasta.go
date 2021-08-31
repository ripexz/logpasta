package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/ripexz/logpasta/clipboard"
)

var version = "v0.3.3"

func main() {
	initLogger()
	conf := loadConfig()
	checkForCommands()

	var content string
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// piped
		bytes, _ := ioutil.ReadAll(os.Stdin)
		content = string(bytes)
	} else {
		content = strings.Join(flag.Args(), " ")
	}

	if content == "" {
		log.Printf("No input detected, see 'logpasta help' for usage")
		os.Exit(0)
	}

	// make request
	var output string
	paste, err := saveLog(conf, content)
	if err != nil {
		conf.Silent = false
		output = fmt.Sprintf("Failed to save log: %s", err.Error())
	} else {
		pasteURL := fmt.Sprintf("%s/paste/%s", conf.BaseURL, paste.UUID)

		output = fmt.Sprintf("Log saved successfully:\n%s", pasteURL)

		if conf.Copy {
			err = clipboard.Copy(pasteURL)
			if err != nil {
				output += fmt.Sprintf("\n(failed to copy to clipboard)")
				if !conf.Silent {
					output += fmt.Sprintf("\nError: %s", err.Error())
				}
			} else {
				output += " (copied to clipboard)"
			}
		}

		if paste.DeleteKey != nil {
			output += fmt.Sprintf(
				"\nYou can delete it early by visiting:\n%s/delete/%s/%s",
				conf.BaseURL, paste.UUID, *paste.DeleteKey,
			)
		}

	}

	if !conf.Silent {
		log.Println(content)
	}

	log.Println(output)
}

func checkForCommands() {
	if len(os.Args) <= 1 {
		return
	}

	switch os.Args[1] {
	case "version":
		printVersion()
		os.Exit(0)
	case "help":
		printUsage()
		os.Exit(0)
	}
}
