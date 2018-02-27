package services

import (
	"encoding/json"
	"fmt"
	"soccergist/implementations/go/dataobject"
	"soccergist/implementations/go/utility"
	"strconv"
)

//HandlePostBackRecieved - function to handle the postbacks
func HandlePostBackRecieved(postback dataobject.PostBack, sender dataobject.Sender) (response string) {
	title := postback.Title
	payload := postback.Payload
	var recipient dataobject.Recipient
	var message dataobject.ResponseMessage
	recipient = dataobject.Recipient{
		ID: sender.ID,
	}

	fmt.Println(payload)

	if payload == "league-table-postback" {
		leagueID := 445
		return loadLeagueTable(leagueID, sender)

	} else {
		message = dataobject.ResponseMessage{
			Text: title + " coming soon",
		}

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

func loadLeagueTable(leagueID int, sender dataobject.Sender) (response string) {
	url := "http://api.football-data.org/v1/competitions/" + strconv.Itoa(leagueID) + "/leagueTable"
	dataResponse := utility.SendGetRequest(url)
	var decodedResponse map[string]interface{}
	var elements []dataobject.Element

	err := json.Unmarshal([]byte(dataResponse), &decodedResponse)
	utility.FailOnError(err, "Cannot Convert this response to a reasonable data")

	//lets get the standing item
	standings := decodedResponse["standing"].([]interface{})

	for i := 0; i < 4; i++ {
		//standing has been converted into a map of string key and empty interface value.
		var buttons []dataobject.Button
		standing := standings[i].(map[string]interface{})
		teamName := standing["teamName"].(string)
		positionResult := standing["position"].(float64)
		playedResult := standing["playedGames"].(float64)
		pointResult := standing["points"].(float64)

		position := strconv.FormatFloat(positionResult, 'f', 0, 64)
		played := strconv.FormatFloat(playedResult, 'f', 0, 64)
		points := strconv.FormatFloat(pointResult, 'f', 0, 64)

		title := "Position" + position + ": " + teamName
		subtitle := "Mathes Played: " + played + "\nPoints: " + points
		imageURL := dataobject.TeamLogos[teamName]

		button := dataobject.Button{
			Title:   "more details",
			Type:    "postback",
			Payload: "league-table-position-" + position + "-mores-detail-postback",
		}

		buttons = append(buttons, button)

		newElement := dataobject.Element{
			Title:    title,
			SubTitle: subtitle,
			ImageURL: imageURL,
			Buttons:  buttons,
		}

		elements = append(elements, newElement)
	}

	listPayload := dataobject.ListPayload{
		TemplateType:    "list",
		TopElementStyle: "compact",
		Elements:        elements,
	}

	attachment := dataobject.Attachment{
		Type:    "template",
		Payload: listPayload,
	}

	message := dataobject.ResponseMessage{
		Attachment: &attachment,
	}

	recipient := dataobject.Recipient{
		ID: sender.ID,
	}

	jsonResponse := dataobject.JSONResponse{
		Recipient: recipient,
		Message:   message,
	}

	b, err := json.Marshal(jsonResponse)
	utility.FailOnError(err, "Cannot generate a json data from this object")
	response = string(b)
	return
}
