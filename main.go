package main

import (
	"github.com/javiergomezve/allergycron/allergy_api"
	"github.com/javiergomezve/allergycron/utils"
	"github.com/joho/godotenv"
	"github.com/robfig/cron"
	"log"
	"os"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("error loading .env")
	}

	location, err := time.LoadLocation(os.Getenv("CRON_TIMEZONE"))
	if err != nil {
		panic(err)
	}

	cronJob := cron.NewWithLocation(location)
	err = cronJob.AddFunc(os.Getenv("CRON_SCHEDULE"), func() {
		dailyAverageMessage, err := allergy_api.GetHourlyLoadData()
		if err != nil {
			panic(err)
		}

		historicalAverageMessage, err := allergy_api.GetCurrentChartData()
		if err != nil {
			panic(err)
		}

		slackMessage := *dailyAverageMessage + "\n" + *historicalAverageMessage
		err = utils.SendSlackMessage(slackMessage)
		if err != nil {
			panic(err)
		}

		log.Println("✅Successfully sent Slack message: " + slackMessage)
	})

	if err != nil {
		log.Println("❌The message could not been sent")
	}
}
