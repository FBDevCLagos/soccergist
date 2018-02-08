package services

import (
	"encoding/json"
	"soccergist/implementations/go/dataobject"
	"soccergist/implementations/go/utility"
)

//HandlePostBackRecieved - function to handle the postbacks
func HandlePostBackRecieved(postback dataobject.PostBack, sender dataobject.Sender) (response string) {
	title := postback.Title
	recipient := dataobject.Recipient{
		ID: sender.ID,
	}
	message := dataobject.ResponseMessage{
		Text: title + " coming soon",
	}
	jsonResponse := dataobject.JSONResponse{
		Recipient: recipient,
		Message:   message,
	}

	b, err := json.Marshal(jsonResponse)
	utility.FailOnError(err, "Cannot Marshall response to json")
	response = string(b)
	return
}
