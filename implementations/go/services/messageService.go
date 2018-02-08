package services

import (
	"encoding/json"
	"soccergist/implementations/go/dataobject"
	"soccergist/implementations/go/utility"
)

//HandleMessageRecieved -function to handle message recieved
func HandleMessageRecieved(message dataobject.Message, sender dataobject.Sender) (response string) {

	//the response is dependent on the type of message we recieve
	//but for now, we assume that all message means greeting
	//and we bombard our users with the options available.
	response = ShowDefaultMenu(sender)
	return
}

//ShowDefaultMenu - function to show the defuault menu with buttons
func ShowDefaultMenu(sender dataobject.Sender) (response string) {
	recipient := dataobject.Recipient{
		ID: sender.ID,
	}

	scheduleButton := dataobject.Button{
		Type:    "postback",
		Title:   "View Match Schedules",
		Payload: "match-schedule-postback",
	}

	highlightButton := dataobject.Button{
		Type:    "postback",
		Title:   "View Match Highlights",
		Payload: "match-highligh-postback",
	}

	tableButton := dataobject.Button{
		Type:    "postback",
		Title:   "View League Table",
		Payload: "league-table-postback",
	}

	buttons := []dataobject.Button{scheduleButton, highlightButton, tableButton}

	responsePayload := dataobject.Payload{
		Text:         "Hi, What do you want to do?",
		TemplateType: "button",
		Buttons:      buttons,
	}

	attachmentResponse := dataobject.Attachment{
		Type:    "template",
		Payload: responsePayload,
	}

	responseMessage := dataobject.ResponseMessage{
		Attachment: attachmentResponse,
	}

	jsonResponse := dataobject.JSONResponse{
		Recipient: recipient,
		Message:   responseMessage,
	}

	b, err := json.Marshal(jsonResponse)
	utility.FailOnError(err, "Cannot Marshall This response accordingly...")
	response = string(b)
	return response
}
