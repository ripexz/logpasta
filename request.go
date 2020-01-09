package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func saveLog(conf *Config, content string) (*Paste, error) {
	client := http.Client{Timeout: time.Second * time.Duration(conf.Timeout)}

	data := PasteData{
		Paste: Paste{
			Content: content,
		},
	}

	payload, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	zipper := gzip.NewWriter(&buf)
	if _, err = zipper.Write(payload); err != nil {
		return nil, err
	}
	if err = zipper.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest(
		http.MethodPost,
		fmt.Sprintf("%s/api/v1/pastes.json", conf.BaseURL),
		bytes.NewReader(buf.Bytes()),
	)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Content-Encoding", "gzip")
	req.Header.Set("User-Agent", fmt.Sprintf(
		"Logpasta-CLI/%s",
		strings.Replace(version, "v", "", -1),
	))

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return nil, fmt.Errorf("failed to make request: %s", res.Status)
	}

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(resData, &data)
	return &data.Paste, err
}
