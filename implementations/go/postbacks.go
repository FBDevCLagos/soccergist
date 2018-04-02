package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/FBDevCLagos/soccergist/implementations/go/data"
)

func handleLeagueTablePostbackEvent(msgEvnt, senderID string) {
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

func handleMatchSchedulesPostbackEvent(payload, senderID string) {
	premierLeague := data.PremierLeagueInfo()
	rr := regexp.MustCompile("(\\d+)")
	matchday := premierLeague.PresentMatchday()
	if mday := rr.FindStringSubmatch(payload); len(mday) > 0 {
		matchday, _ = strconv.Atoi(mday[0])
	}

	fixtures := premierLeague.GetMatchdayFixtures(matchday)
	msg := fmt.Sprintf("Fixtures for Matchday: %d", matchday)

	if fixtures == nil {
		return
	}

	for _, fixture := range fixtures.Fixtures {
		if fixture.Status == "FINISHED" {
			msg = fmt.Sprintf("%s\n-----\n%s VS %s => (%d : %d)", msg, fixture.HomeTeamName, fixture.AwayTeamName, fixture.Result.GoalsHomeTeam, fixture.Result.GoalsAwayTeam)
		} else if t, err := time.Parse(time.RFC3339, fixture.Date); err == nil {
			msg = fmt.Sprintf("%s\n-----\n%s VS %s - %s", msg, fixture.HomeTeamName, fixture.AwayTeamName, t.Format("Mon, Jan 2, 3:04PM"))
		} else {
			msg = fmt.Sprintf("%s\n-----\n%s VS %s - %s", msg, fixture.HomeTeamName, fixture.AwayTeamName, fixture.Status)
		}
	}
	reply := buildTextMsg(senderID, msg)

	sendResponse(reply)
	sendMatchFixuresPagination(senderID, matchday, premierLeague.PresentMatchday(), premierLeague.TotalMatchdays())
}

func sendMatchFixuresPagination(senderID string, matchday, currentMatchday, totalMatchdays int) {
	contents := []quickReply{}

	if day := matchday - 2; day > 0 {
		content := quickReply{
			ContentType: "text",
			Payload:     fmt.Sprintf("match-schedules-postback-%d", day),
			Title:       fmt.Sprintf("<< matchday %d", day),
		}
		contents = append(contents, content)
	}

	if day := matchday - 1; day > 0 {
		content := quickReply{
			ContentType: "text",
			Payload:     fmt.Sprintf("match-schedules-postback-%d", day),
			Title:       fmt.Sprintf("< matchday %d", day),
		}
		contents = append(contents, content)
	}

	if matchday != currentMatchday {
		content := quickReply{
			ContentType: "text",
			Payload:     fmt.Sprintf("match-schedules-postback-%d", currentMatchday),
			Title:       "current matchday",
		}
		contents = append(contents, content)
	}

	if day := matchday + 1; day <= totalMatchdays {
		content := quickReply{
			ContentType: "text",
			Payload:     fmt.Sprintf("match-schedules-postback-%d", day),
			Title:       fmt.Sprintf("matchday %d >", day),
		}
		contents = append(contents, content)
	}

	if day := matchday + 2; day <= totalMatchdays {
		content := quickReply{
			ContentType: "text",
			Payload:     fmt.Sprintf("match-schedules-postback-%d", day),
			Title:       fmt.Sprintf("matchday %d >>", day),
		}
		contents = append(contents, content)
	}

	reply := buildQuickReply("navigation", senderID, contents)
	sendResponse(reply)
}

func handleLeagueMoreHighlightsPostbackEvent(payload, senderID string) {
	handleLeagueHighlights(payload, senderID)
}

func handleLeagueHighlightsPostbackEvent(payload, senderID string) {
	handleLeagueHighlights("", senderID)
}

func handleLeagueHighlights(payload, senderID string) {
	elements := []element{}
	posts := data.Highlights(payload)

	for _, post := range posts {
		// TODO: handle more than one highlight url
		btn := button{Type: "web_url", Title: "Watch highlight", URL: post.URLs[0]}
		element := buildBasicElement(post.Title, "", "https://i.vimeocdn.com/portrait/6640852_640x640")
		element.Buttons = []button{btn}
		elements = append(elements, element)

		if len(elements) == 9 {
			break
		}
	}

	if len(posts) > 9 {
		viewMore := buildBasicElement("View More Highlights", "", "http://icons-for-free.com/free-icons/png/512/1814113.png")
		btn := buildPostbackBtn("More Highlights", posts[9].Name)
		viewMore.Buttons = []button{btn}

		elements = append(elements, viewMore)
	}

	reply := templateResponse{}
	reply.Recipient.ID = senderID
	reply.Message.Attachment.Type = "template"
	reply.Message.Attachment.Payload.TemplateType = "generic"
	reply.Message.Attachment.Payload.Elements = elements

	sendResponse(reply)
}
