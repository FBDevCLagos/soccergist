package main

import (
	"fmt"

	"github.com/FBDevCLagos/soccergist/implementations/go/data"
)

func handleLeagueTablePostbackEvent(msgEvnt postbackEvent, senderID string) {
	league := data.PremierLeagueInfo()
	leagueTable := league.Table()
	firstFour := data.FirstFour(leagueTable)
	elements := []element{}
	moreDetailsBtn := buildPostbackBtn("more details", "")
	viewMoreBtn := buildPostbackBtn("view more", "league-table-view-more-postback")

	for _, team := range firstFour {
		moreDetailsBtn.Payload = fmt.Sprintf("league-table-position-%d-more-details-postback", team.Position)
		element := buildBasicElement(
			fmt.Sprintf("Position %d: %s", team.Position, team.Name),
			fmt.Sprintf("Matches played: %d \n Points: %d", team.MatchPlayed, team.Points),
			team.Crest,
		)

		element.Buttons = append(element.Buttons, moreDetailsBtn)
		elements = append(elements, element)
	}

	reply := templateResponse{}
	reply.Recipient.ID = senderID
	reply.Message.Attachment.Type = "template"
	reply.Message.Attachment.Payload.TemplateType = "list"
	reply.Message.Attachment.Payload.TopElementStyle = "large"
	reply.Message.Attachment.Payload.Elements = elements
	reply.Message.Attachment.Payload.Buttons = []button{viewMoreBtn}

	sendResponse(reply)
}
