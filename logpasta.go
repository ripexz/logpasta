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
	"time"
)

type Paste struct {
	ID      int64  `json:"id,omitempty"`
	Content string `json:"content,omitempty"`
}

type PasteData struct {
	Paste `json:"paste"`
}

var (
	silent bool
)

func main() {
	// parse flags
	flag.BoolVar(&silent, "s", false, "silent mode")

	var content string
	fi, _ := os.Stdin.Stat()
	if (fi.Mode() & os.ModeCharDevice) == 0 {
		// piped
		bytes, _ := ioutil.ReadAll(os.Stdin)
		content = string(bytes)
	} else {
		content = flag.Arg(0)
	}

	// make request
	var output string
	id, err := saveLog(content)
	if err != nil {
		silent = false
		output = fmt.Sprintf("Failed to save log: %s", err.Error())
	} else {
		output = fmt.Sprintf("Log saved successfully: http://localhost:9999/api/v1/pastes/%d.json", id)
	}

	if !silent {
		log.Println(content)
	}

	log.Println(output)
}

func saveLog(content string) (int64, error) {
	client := http.Client{Timeout: time.Second}

	data := &PasteData{
		Paste: Paste{
			Content: content,
		},
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	var buf bytes.Buffer
	zipper := gzip.NewWriter(&buf)
	if _, err = zipper.Write(payload); err != nil {
		return 0, err
	}
	if err = zipper.Close(); err != nil {
		return 0, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		"http://localhost:9999/api/v1/pastes.json",
		bytes.NewReader(buf.Bytes()),
	)
	if err != nil {
		return 0, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")

	res, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return 0, fmt.Errorf("failed to make request: %s", res.Status)
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(resData, data)
	return data.Paste.ID, err
}
