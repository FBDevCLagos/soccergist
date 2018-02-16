package utils

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func APIRequest(url string, requestMethod string, payload interface{}) (*http.Response, error) {
	var req *http.Request
	var err error

	if payload != nil {
		// Parse the response payload
		pkg, err := json.Marshal(payload)
		if err != nil {
			log.Println("Sending response parsing in an error: ", err)
			return nil, err
		}
		body := bytes.NewBuffer(pkg)
		req, err = http.NewRequest(requestMethod, url, body)
	} else {
		req, err = http.NewRequest(requestMethod, url, nil)
	}

	if err != nil {
		log.Println("Error creating request: ", err)
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}

	return client.Do(req)
}
