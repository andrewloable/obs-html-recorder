package main

import (
	"flag"
	"log"
	"time"

	"github.com/andrewloable/obs-html-recorder/config"
	"github.com/andrewloable/obs-html-recorder/obs"
	"github.com/andrewloable/obs-html-recorder/profile"
)

func main() {

	widthFlag := flag.Int("w", 1920, "width of the browser")
	heightFlag := flag.Int("h", 1080, "height of the browser")
	urlFlag := flag.String("url", "https://fast.com", "the url of the html to be recorded")
	secondsFlag := flag.Int("s", 10, "time in seconds to record the html")
	flag.Parse()

	_, err := config.ReadConfig()
	if err != nil {
		log.Println(err)
		return
	}
	if config.AppConfig.IsReady {
		log.Println(config.AppConfig)
	}

	prof := profile.Profile{
		Width:  *widthFlag,
		Height: *heightFlag,
	}

	log.Println("set profile ", prof)
	client, err := obs.InitiateObsRecorder(prof)
	if err != nil {
		log.Println("initiation error ", err)
	}

	log.Println("record html")
	obs.RecordHTML(client, *urlFlag, *secondsFlag, prof)
	time.Sleep(time.Second * 2)
	client.Disconnect()
	obs.TerminateObs()
}
