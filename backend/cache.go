package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func GetCache(hash string) (response Response, exists bool) {
	url := fmt.Sprintf("https://kvdb.io/7XHQsAyLHuFZFNSWYksVGC/%s", hash)
	log.Printf("[%s] Getting %s", hash, url)

	res, err := http.Get(url)
	if res.StatusCode == 404 {
		log.Printf("[%s] Not Found", hash)
		return
	}
	if err != nil {
		return
	}

	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return
	}

	json.Unmarshal(content, &response)

	log.Printf("[%s] Got %v", hash, response)

	return response, true
}

func SetCache(hash string, response Response) error {
	url := fmt.Sprintf("https://kvdb.io/7XHQsAyLHuFZFNSWYksVGC/%s", hash)
	log.Printf("Setting %s", url)

	content, err := json.Marshal(response)
	if err != nil {
		log.Printf("%v", err)
		return err
	}

	if _, err = http.Post(url, "application/json", bytes.NewBuffer(content)); err != nil {
		log.Printf("%v", err)
	}

	log.Printf("[%s] Set %v", hash, response)

	return err
}
