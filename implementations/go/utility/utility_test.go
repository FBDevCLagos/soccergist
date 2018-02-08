package utility

import (
	"log"
	"testing"
)

func TestGetSecretKey(t *testing.T) {
	response := GetSecretKey("../.env")

	log.Fatal("Response is: ", response)
}
