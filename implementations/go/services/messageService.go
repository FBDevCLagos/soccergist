package services

import (
	"encoding/json"
	"soccergist/implementations/go/dataobject"
	"soccergist/implementations/go/utility"
	"strconv"
)

//HandleMessageRecieved -function to handle message recieved
func HandleMessageRecieved(message dataobject.Message, sender dataobject.Sender) (response string) {

	/**
	**/

	if message.QuickReply.Payload != "" {
		leagueID := 445
		matchDay, err := strconv.Atoi(message.QuickReply.Payload)
		utility.FailOnError(err, "Invalid Matchday Number provided")
		response = LoadMatchScheduleReply(sender, leagueID, matchDay)
	} else {
		response = ShowDefaultMenu(sender)
	}

	// response = ShowQuickReplies(sender, message.QuickReply)
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

	responsePayload := dataobject.ButtonPayload{
		Text:         "Hi, What do you want to do? Go Implentation",
		TemplateType: "button",
		Buttons:      buttons,
	}

	attachmentResponse := dataobject.Attachment{
		Type:    "template",
		Payload: &responsePayload,
	}

	responseMessage := dataobject.ResponseMessage{
		Attachment: &attachmentResponse,
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

//ShowQuickReplies - function to demonstrate quick reply Demo
func ShowQuickReplies(sender dataobject.Sender, quickReplyPayload dataobject.MessageQuickReply) (response string) {
	var quickReplies []dataobject.QuickReply
	start := 1
	end := 5
	limit := 20
	var payloadValue int
	if quickReplyPayload.Payload != "" {
		payloadValue, _ = strconv.Atoi(quickReplyPayload.Payload)
	} else {
		payloadValue = 1
	}

	newStartPoint := payloadValue - 2
	if newStartPoint > 0 {
		start = newStartPoint
	}

	newEndPoint := payloadValue + 2
	if newEndPoint > limit {
		end = limit
		start = limit - 4
	} else {
		if newEndPoint > end {
			end = newEndPoint
		}
	}

	for i := start; i <= end; i++ {
		if i == payloadValue {
			continue
		}
		strContent := strconv.Itoa(i)
		quickReplies = append(quickReplies, dataobject.QuickReply{
			ContentType: "text",
			Title:       strContent,
			Payload:     strContent,
		})
	}

	quickReplyMessage := dataobject.QuickResponseMessage{
		Text:         "Naviagation Options",
		QuickReplies: quickReplies,
	}

	recipient := dataobject.Recipient{
		ID: sender.ID,
	}

	jsonOutputMessage := dataobject.JSONResponse{
		Recipient: recipient,
		Message:   quickReplyMessage,
	}

	b, err := json.Marshal(jsonOutputMessage)
	utility.FailOnError(err, "Cannot Marshal This message to json")

	response = string(b)
	return
}
