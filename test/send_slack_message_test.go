package test

import (
	"github.com/javiergomezve/allergycron/utils"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestSendSlackMessage(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("error loading .env")
	}

	err = utils.SendSlackMessage("Test message!")
	if err != nil {
		t.Errorf("error sending slack message: %s", err)
	}
}
