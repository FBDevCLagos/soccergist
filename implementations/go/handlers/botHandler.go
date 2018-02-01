package handlers

import (
	"fmt"
	"net/http"
)

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
