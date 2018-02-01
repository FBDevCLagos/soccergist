package handlers

import (
	"fmt"
	"net/http"
)

//HomeHandler - function to handle lading page routes
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "welcome to the movie guide bot challenge")
}
