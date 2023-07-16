package test

import (
	"github.com/javiergomezve/allergycron/allergy_api"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestAllergyApi(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println("error loading .env")
	}

	message, err := allergy_api.GetHourlyLoadData()
	if err != nil {
		t.Errorf("error getting hourly load data: %s", err)
	}

	if message == nil {
		t.Errorf("error getting hourly load data: message nil")
	}

	if *message == "" {
		t.Errorf("error getting hourly load data: message is empty")
	}

	message, err = allergy_api.GetCurrentChartData()
	if err != nil {
		t.Errorf("error getting current chart data: %s", err)
	}

	if message == nil {
		t.Errorf("error getting current chart data: message nil")
	}

	if *message == "" {
		t.Errorf("error getting current chart data: message is empty")
	}
}
