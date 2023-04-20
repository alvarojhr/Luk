package main

import (
	"archive/tar"
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"sync"
	"strings"
	"bytes"
	"runtime"
)

const maxBufferSize = 1000 * 1024 * 1024

type FileTask struct {
	Header  *tar.Header
	Content io.Reader
}

var messagePool = &sync.Pool{
	New: func() interface{} {
		return &Message{}
	},
}


func readAndProcessTarGz(filename string, messages chan<- *Message) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	gzipReader, err := gzip.NewReader(file)
	if err != nil {
		return err
	}
	defer gzipReader.Close()

	tarReader := tar.NewReader(gzipReader)
	qnMessages := 0

	var wg sync.WaitGroup
	fileChan := make(chan *FileTask)

	workerCount := runtime.NumCPU()
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for fileTask := range fileChan {
				processFile(fileTask.Header, fileTask.Content,messages)
				qnMessages++
			}
		}()
	}

	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		content := new(bytes.Buffer)
		if _, err := io.Copy(content, tarReader); err != nil {
			return err
		}

		fileTask := &FileTask{
			Header:  header,
			Content: content,
		}

		fileChan <- fileTask
	}

	close(fileChan)
	wg.Wait()
	fmt.Printf("Total files: %d",qnMessages)
	return nil
}

// func processFolderContents(header *tar.Header) {
// 	fmt.Printf("Processing folder: %s\n", header.Name)
// }

func processFile(header *tar.Header, reader io.Reader, messages chan<- *Message) {
	//fmt.Printf("Processing file: %s\n", header.Name)

	scanner := bufio.NewScanner(reader)
	buf := make([]byte, 0, bufio.MaxScanTokenSize)
	scanner.Buffer(buf, maxBufferSize)
	var messageBuilder MessageBuilder

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		messageBuilder.ParseLine(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file content: %v\n", err)
		return
	}

	//message := 
	message := messageBuilder.Build()
	messages <- &message
	//fmt.Printf("Procesed message: %s\n", message.MessageID)

	// Do something with the parsed message
	messageBuilder.Reset() // Reset the builder for reuse
}
