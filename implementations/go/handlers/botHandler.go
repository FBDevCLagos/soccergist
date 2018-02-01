package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/LordRahl90/soccergist/implementations/go/utility"
)

//JSONRequest struct
type JSONRequest struct {
	Object string      `json:"object"`
	Entry  []EntryItem `json:"entry"`
}

//JSONResponse - struct
type JSONResponse struct {
	Recipient Recipient       `json:"recipient"`
	Message   ResponseMessage `json:"message"`
}

//ResponseMessage struct
type ResponseMessage struct {
	Text string `json:"text"`
}

//EntryItem Struct
type EntryItem struct {
	ID        string          `json:"id"`
	Time      int64           `json:"time"`
	Messaging []MessagingItem `json:"messaging"`
}

//MessagingItem struct for handling the messaging values
type MessagingItem struct {
	Sender    Sender    `json:"sender"`
	Recipient Recipient `json:"recipient"`
	Timestamp int64     `json:"timestamp"`
	Message   Message   `json:"message"`
}

//Message struct handles the message that is being sent across.
type Message struct {
	MID  string `json:"mid"`  //message ID
	Seq  int    `json:"seq"`  //sequence
	Text string `json:"text"` //content of the message
}

//Sender struct - To Manage the sender
type Sender struct {
	ID string `json:"id"`
}

//Recipient struct to manage the recipient
type Recipient struct {
	ID string `json:"id"`
}

//WebHookHandler - function to handler webhooks
func WebHookHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()

	hubMode := queryString.Get("hub.mode")
	hubChallenge := queryString.Get("hub.challenge")
	hubVerifyToken := queryString.Get("hub.verify_token")

	fmt.Println(hubVerifyToken)

	if hubMode != "subscribe" {
		w.WriteHeader(403)
		fmt.Fprint(w, "Invalid Mode Discovered!")
		return
	}

	if hubVerifyToken != "only the strong will continue" {
		w.WriteHeader(403)
		fmt.Fprint(w, "Invalid/Unhandled token detected")
		return
	}

	fmt.Fprint(w, hubChallenge)
}

//WebHookPostHandler - function to handle webhook post requests
func WebHookPostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Body)
	var jsonRequest JSONRequest
	err := json.NewDecoder(r.Body).Decode(&jsonRequest)
	utility.FailOnError(err, "Cannot decode this request data successfully!!")

	entryItem := jsonRequest.Entry[0]
	messaging := entryItem.Messaging[0]

	senderID := messaging.Sender.ID
	message := messaging.Message.Text

	if senderID == "" {
		fmt.Fprint(w, utility.ReturnErrorMessage("Sender ID not found", "Sender Not Found"))
		return
	}

	if message == "" {
		fmt.Fprint(w, utility.ReturnErrorMessage("Message is Empty", "No Message Found"))
		return
	}

	responseMessage := ResponseMessage{
		Text: "I recieved your message (" + message + ") and i have sent it to my Oga at the top: LordRahl",
	}

	recipient := Recipient{
		ID: senderID,
	}

	jsonResponse := JSONResponse{
		Recipient: recipient,
		Message:   responseMessage,
	}

	b, err := json.Marshal(jsonResponse)
	utility.FailOnError(err, "Cannot Convert this to JSON")
	responseString := string(b)

	fmt.Println(responseString)

	//sending the post request
	accessToken := "EAAEDNuZAnTygBAMDqKr4H3eJQQnZBhyi25gbotURZBzOSO0nhzkhSdkIf2vCiA2sM9u4L3h2Hfd6MgY2iKr70SLWoeWnhuQ8RcSZBzjXSjBfaqi4tufTO20Dq1PMDZC1kHlhR8mYAH6pZCL7F122HKUEwgZA7DbQeU7oHqE3D43mwZDZD"
	endPoint := "https://graph.facebook.com/v2.6/me/messages?access_token=" + accessToken
	request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer([]byte(responseString))) //creates a request
	request.Header.Set("Content-Type", "application/json")
	utility.FailOnError(err, "Cannot Complete this request")

	client := utility.GetHTTPClient()
	response, err := client.Do(request) //sends the request to the desired endpoint and keeps the response
	utility.FailOnError(err, "Cannot Process this request")
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body) //gets the body of the response
	fmt.Println(string(body))

	w.Header().Add("COntent-Type", "application/json")
	fmt.Fprint(w, responseString) //prints the good news to the user
}
