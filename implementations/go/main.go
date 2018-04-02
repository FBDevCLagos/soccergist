package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"strings"

	"github.com/FBDevCLagos/soccergist/implementations/go/utils"

	"github.com/julienschmidt/httprouter"
)

// VerificationToken is the random string entered in the verification prompt when setting up the app on Facbook
// It can be any string provide it matches what you will enter in the setup prompt
const VerificationToken = "bots are awesome"

var (
	AccessToken = os.Getenv("ACCESS_TOKEN")
	fbURL       = fmt.Sprintf("https://graph.facebook.com/v2.6/me/messages?access_token=%s", AccessToken)
)

func verifyWebhook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	if mode == "subscribe" && token == VerificationToken {
		fmt.Fprint(w, challenge)
	} else {
		w.WriteHeader(http.StatusForbidden)
	}
}

func handleWebhookEvents(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Parse the request payload
	payload := webhookPayload{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		log.Println("Unmarshalling webhook payload resulted in an error: ", err)
		return
	}

	// Make sure this is a page subscription
	if payload.Object == "page" {
		// Iterate over each entry
		// There may be multiple if batched
		for _, entry := range payload.Entry {
			// Iterate over each messaging event
			for _, messaging := range entry.Messaging {
				switch {
				case !reflect.DeepEqual(messaging.Message, messageEvent{}):
					handleMessageEvent(messaging.Message, messaging.Sender.ID)
				case !reflect.DeepEqual(messaging.Postback, postbackEvent{}):
					handlePostbackEvent(messaging.Postback, messaging.Sender.ID)
				default:
					log.Printf("No handler found for: %+v\n", messaging.Message)
				}
			}
		}
	}
}

var postbackHandlers = map[string]func(string, string){
	"league-table-postback":      handleLeagueTablePostbackEvent,
	"match-schedules-postback":   handleMatchSchedulesPostbackEvent,
	"league-highlights-postback": handleLeagueHighlightsPostbackEvent,

	"More Highlights": handleLeagueMoreHighlightsPostbackEvent,
}

func handlePostbackEvent(msgEvnt postbackEvent, senderID string) {

	if postbackHandler, ok := postbackHandlers[msgEvnt.Payload]; ok {
		postbackHandler(msgEvnt.Payload, senderID)
		return
	} else if postbackHandler, ok := postbackHandlers[msgEvnt.Title]; ok {
		postbackHandler(msgEvnt.Payload, senderID)
		return
	}

	reply := templateResponse{}
	reply.Recipient.ID = senderID
	reply.Message.Text = fmt.Sprintf("%s - coming soon 🤠", msgEvnt.Title)
	sendResponse(reply)
}

func handleMessageEvent(msgEvnt messageEvent, senderID string) {
	if msgEvnt.QuickReply.Payload != "" {
		handleQuickReplyEvent(msgEvnt, senderID)
		return
	}
	handleTextMessageEvent(msgEvnt, senderID)
}

func handleTextMessageEvent(msgEvnt messageEvent, senderID string) {
	reply := templateResponse{}
	reply.Recipient.ID = senderID
	reply.Message.Attachment.Type = "template"
	reply.Message.Attachment.Payload.TemplateType = "button"
	reply.Message.Attachment.Payload.Text = "What do you want to do?"

	matchSchedulesPostbackBtn := buildPostbackBtn("View match schedules", "match-schedules-postback")
	leagueTablePostbackBtn := buildPostbackBtn("View league table", "league-table-postback")
	leagueHighlightsBtn := buildPostbackBtn("View Highlights", "league-highlights-postback")

	reply.Message.Attachment.Payload.Buttons = []button{matchSchedulesPostbackBtn, leagueHighlightsBtn, leagueTablePostbackBtn}
	sendResponse(reply)
}

func handleQuickReplyEvent(msgEvnt messageEvent, senderID string) {
	payload := msgEvnt.QuickReply.Payload
	if strings.Contains(payload, "match-schedules-postback-") {
		handleMatchSchedulesPostbackEvent(payload, senderID)
		return
	}
	log.Println("Unrecognized payload")
}

func sendResponse(payload interface{}) {
	utils.APIRequest(fbURL, "POST", payload)
}

func setupRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/webhook", verifyWebhook)
	r.POST("/webhook", handleWebhookEvents)
	return r
}

func main() {
	err := http.ListenAndServe(":3000", setupRouter())
	if err != nil {
		panic(err)
	}
}
