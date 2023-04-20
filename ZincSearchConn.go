package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

const (
	zincAPIURL     = "https://api.zinc.dev/api/alvaro_organization_826/stream1/_multi"
	zincAPIUser    = "alvarojhr96@gmail.com"
	zincAPIKey     = "45vso9ibE086kL1p2Z73"
	batchSize      = 1000
	concurrentWorkers = 5
)

func uploadWorker(messages <-chan *Message, wg *sync.WaitGroup) {
	defer wg.Done()

	batch := make([]*Message, 0, batchSize)

	for msg := range messages {
		batch = append(batch, msg)

		if len(batch) >= batchSize {
			if err := sendBatchToZincSearch(batch); err != nil {
				// Handle error
			}
			batch = batch[:0]
		}
	}

	if len(batch) > 0 {
		if err := sendBatchToZincSearch(batch); err != nil {
			// Handle error
		}
	}
}

func sendBatchToZincSearch(batch []*Message) error {
	var buf bytes.Buffer

	for _, msg := range batch {
		jsonData, err := json.Marshal(msg)
		if err != nil {
			return err
		}
		buf.Write(jsonData)
		buf.WriteString("\n")
	}

	req, err := http.NewRequest("POST", zincAPIURL, strings.NewReader(buf.String()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-ndjson")
	req.SetBasicAuth(zincAPIUser, zincAPIKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("ZincSearch API returned status code %d: %s", resp.StatusCode, string(body))
	}

	return nil
}
