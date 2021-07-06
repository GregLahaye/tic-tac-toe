package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func GetCache(hash string) (response Response, exists bool) {
	url := fmt.Sprintf("https://kvdb.io/9zfvdU7e32RxRyMStNm2zP/%s", hash)

	res, err := http.Get(url)
	if res.StatusCode == 404 {
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

	return response, true
}

func SetCache(hash string, response Response) error {
	url := fmt.Sprintf("https://kvdb.io/9zfvdU7e32RxRyMStNm2zP/%s", hash)

	content, err := json.Marshal(response)
	if err != nil {
		return err
	}

	_, err = http.Post(url, "application/json", bytes.NewBuffer(content))

	return err
}
