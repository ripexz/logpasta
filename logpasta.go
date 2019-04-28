package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	var silent bool
	flag.BoolVar(&silent, "s", false, "silent mode")

	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// piped
		bytes, _ := ioutil.ReadAll(os.Stdin)
		str := string(bytes)
		if !silent {
			log.Println(str)
		}
	} else {
		if !silent {
			log.Println(flag.Args())
		}
	}
}
