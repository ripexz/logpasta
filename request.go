package main

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func saveLog(conf *Config, content string) (string, error) {
	client := http.Client{Timeout: time.Second * time.Duration(conf.Timeout)}

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
	req.Header.Set("User-Agent", fmt.Sprintf("Logpasta CLI %s", version))

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
	return fmt.Sprintf("%s/paste/%s", conf.BaseURL, data.Paste.UUID), err
}
