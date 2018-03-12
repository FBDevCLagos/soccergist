package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"soccergist/implementations/go/dataobject"
	"soccergist/implementations/go/utility"
	"time"

	"soccergist/implementations/go/services"
)

//WebHookHandler - function to handler webhooks
func WebHookHandler(w http.ResponseWriter, r *http.Request) {
	queryString := r.URL.Query()
	fmt.Println(queryString)

	hubMode := queryString.Get("hub.mode")
	hubChallenge := queryString.Get("hub.challenge")
	hubVerifyToken := queryString.Get("hub.verify_token")

	if hubVerifyToken == "" {
		fmt.Fprint(w, "Please provide a valid hub verification token")
		return
	}

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
	var jsonRequest dataobject.JSONRequest
	err := json.NewDecoder(r.Body).Decode(&jsonRequest)
	utility.FailOnError(err, "Cannot decode this request data successfully!!")

	entryItem := jsonRequest.Entry[0]
	messaging := entryItem.Messaging[0]

	sender := messaging.Sender
	message := messaging.Message
	readMessage := messaging.Read
	postBack := messaging.PostBack
	delivery := messaging.Delivery

	var result string

	if message.Text != "" {
		// result = "Processing Message Recieved"
		fmt.Println("Processing Message recieved")
		result = services.HandleMessageRecieved(message, sender)

	} else if readMessage.Watermark != 0 {
		result = "Processing Read Message"
		fmt.Println(result)
		fmt.Fprint(w, result)
		return
	} else if delivery.Watermark != 0 {
		result = "Processing Delivered Message"
		fmt.Println(result)
		fmt.Fprint(w, result)
		return
	} else if postBack.Payload != "" {
		result = services.HandlePostBackRecieved(postBack, sender)
	}

	response := utility.SendPostRequest(result)
	// response := result

	fmt.Println("Response Delivered at " + time.Now().String())
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, response)

	// fmt.Println(jsonRequest)

	// if senderID == "" {
	// 	fmt.Fprint(w, utility.ReturnErrorMessage("Sender ID not found", "Sender Not Found"))
	// 	return
	// }

	// responseMessage := ResponseMessage{
	// 	Text: "I recieved your message (" + message + ") and i have sent it to my Oga at the top: LordRahl",
	// }

	// recipient := Recipient{
	// 	ID: senderID,
	// }

	// jsonResponse := JSONResponse{
	// 	Recipient: recipient,
	// 	Message:   responseMessage,
	// }

	// b, err := json.Marshal(jsonResponse)
	// utility.FailOnError(err, "Cannot Convert this to JSON")
	// responseString := string(b)

	// //sending the post request
	// accessToken := utility.GetSecretKey()
	// endPoint := "https://graph.facebook.com/v2.6/me/messages?access_token=" + accessToken
	// request, err := http.NewRequest("POST", endPoint, bytes.NewBuffer([]byte(responseString))) //creates a request
	// request.Header.Set("Content-Type", "application/json")
	// utility.FailOnError(err, "Cannot Complete this request")

	// fmt.Println("Going to Facebok")

	// client := utility.GetHTTPClient()
	// response, err := client.Do(request) //sends the request to the desired endpoint and keeps the response
	// utility.FailOnError(err, "Cannot Process this request")
	// defer response.Body.Close()

	// body, _ := ioutil.ReadAll(response.Body) //gets the body of the response
	// fmt.Println(string(body))

	// w.Header().Add("Content-Type", "application/json")
	// fmt.Fprint(w, responseString) //prints the good news to the user
}
