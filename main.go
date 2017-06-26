package main

import (
	"log"
	"os"
)

func main() {
	app := NewSlackbot(os.Getenv("SLACK_TOKEN"))

	/* optional shutdown command */
	app.ShutdownCMD = "__shutdown"

	go app.HandleIncomingEvents()

LOOP:
	for {
		select {
		case <-app.Shutdown:
			app.Session.Disconnect()
			break LOOP
		}
	}

	log.Println("finished")
}
