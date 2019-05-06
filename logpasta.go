package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	baseURL = "https://www.logpasta.com"
	version = "v0.1.0"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "version" {
		log.Printf("Logpasta CLI %s", version)
		return
	}

	log.SetPrefix("[logpasta] ")

	// load config from env then from flags
	conf := Config{
		BaseURL: baseURL,
		Silent:  true,
	}
	loadEnv(&conf)
	loadFlags(&conf)

	if conf.Debug {
		log.Printf("Running with config:\n - BaseURL: %s\n - Silent: %v\n - Debug: %v",
			conf.BaseURL,
			conf.Silent,
			conf.Debug,
		)
	}

	var content string
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// piped
		bytes, _ := ioutil.ReadAll(os.Stdin)
		content = string(bytes)
	} else {
		content = strings.Join(flag.Args(), " ")
	}

	// make request
	var output string
	uuid, err := saveLog(&conf, content)
	if err != nil {
		conf.Silent = false
		output = fmt.Sprintf("Failed to save log: %s", err.Error())
	} else {
		output = fmt.Sprintf("Log saved successfully:\n%s/paste/%s", conf.BaseURL, uuid)
	}

	if !conf.Silent {
		log.Println(content)
	}

	log.Println(output)
}

func saveLog(conf *Config, content string) (string, error) {
	client := http.Client{Timeout: time.Second}

	data := &PasteData{
		Paste: Paste{
			Content: content,
		},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	zipper := gzip.NewWriter(&buf)
	if _, err = zipper.Write(payload); err != nil {
		return "", err
	}
	if err = zipper.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/pastes.json", conf.BaseURL),
		bytes.NewReader(buf.Bytes()),
	)
	if err != nil {
		return "", err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return "", fmt.Errorf("failed to make request: %s", res.Status)
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	err = json.Unmarshal(resData, data)
	return data.Paste.UUID, err
}
