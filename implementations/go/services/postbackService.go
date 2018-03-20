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
	} else if payload == "match-schedule-postback" {
		leagueID := 445
		matchDay := 10
		return LoadMatchScheduleReply(sender, leagueID, matchDay)
	}

	message = dataobject.ResponseMessage{
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

//LoadMatchScheduleReply - response to loading match schedule
func LoadMatchScheduleReply(sender dataobject.Sender, leagueID int, matchDay int) (response string) {
	// quickReplyPayload dataobject.MessageQuickReply
	var quickReplies []dataobject.QuickReply
	start := 1
	end := 5
	limit := 38
	var payloadValue int
	payloadValue = matchDay

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
		Text:         loadMatchSchedule(leagueID, matchDay),
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

func loadMatchSchedule(leagueID int, matchDay int) (response string) {
	url := "http://api.football-data.org/v1/competitions/" + strconv.Itoa(leagueID) + "/fixtures?matchday=" + strconv.Itoa(matchDay)
	dataResponse := utility.SendGetRequest(url)
	var decodedFixtures dataobject.Fixtures

	err := json.Unmarshal([]byte(dataResponse), &decodedFixtures)
	utility.FailOnError(err, "Cannont Convert this to json successfully")
	output := `Fixures For Matchday ` + strconv.Itoa(matchDay) + "\n-----------------\n"

	for _, fixture := range decodedFixtures.Fixture {
		output = output + fixture.HomeTeam + " VS " + fixture.AwayTeam + "=> [" +
			strconv.Itoa(fixture.Result.GoalsHomeTeam) + ":" + strconv.Itoa(fixture.Result.GoalsAwayTeam) + "]. Half Time:  (" + strconv.Itoa(fixture.Result.HalfTime.GoalsHomeTeam) +
			":" + strconv.Itoa(fixture.Result.HalfTime.GoalsAwayTeam) + ")\n-----------------------------\n"
	}

	fmt.Println(output)
	return output
}

//Loads the first four teams on the table.
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
