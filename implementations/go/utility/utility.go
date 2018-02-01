package utility

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//Response Struct
type Response struct {
	Status      string      `json:"status"`
	Description string      `json:"description"`
	Message     string      `json:"message"`
	Data        interface{} `json:"data"`
}

//ReturnErrorMessage function
func ReturnErrorMessage(message, description string) (response string) {
	resp := Response{
		Status:      "error",
		Description: description,
		Message:     message,
		Data:        nil,
	}

	b, err := json.Marshal(resp)
	FailOnError(err, "Cannot MArshal to JSON")

	response = string(b)
	return
}

//ReturnErrorMessageWithData function
func ReturnErrorMessageWithData(message, description string, data interface{}) (response string) {
	resp := Response{
		Status:      "error",
		Description: description,
		Message:     message,
		Data:        data,
	}

	b, err := json.Marshal(resp)
	FailOnError(err, "Cannot MArshal to JSON")

	response = string(b)
	return
}

//ReturnSuccessMessage function
func ReturnSuccessMessage(message, description string, data interface{}) (response string) {
	resp := Response{
		Status:      "success",
		Description: description,
		Message:     message,
		Data:        data,
	}

	b, err := json.Marshal(resp)
	FailOnError(err, "Cannot MArshal to JSON")

	response = string(b)
	return
}

//GetHTTPClient - Returns customized HTTP client
func GetHTTPClient() *http.Client {
	client := &http.Client{
		Timeout: time.Second * 30,
	}

	return client
}

//FailOnError function to handle the error logic
func FailOnError(err error, msg string) {
	if err != nil {
		log.Println(msg)
		log.Fatal(err)
	}
}
