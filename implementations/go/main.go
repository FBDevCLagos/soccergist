package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// VerificationToken is the random string entered in the verification prompt when setting up the app on Facbook
// It can be any string provide it matches what you will enter in the setup prompt
const VerificationToken = "bots are awesome"

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

func setupRouter() *httprouter.Router {
	r := httprouter.New()
	r.GET("/webhook", verifyWebhook)
	return r
}

func main() {
	err := http.ListenAndServe(":3000", setupRouter())
	if err != nil {
		panic(err)
	}
}
