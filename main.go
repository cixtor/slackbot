package main

import (
	"log"
	"os"
)

func main() {
	app := NewSlackbot(os.Getenv("SLACK_TOKEN"), os.Getenv("SLACK_COOKIE"))

	/* optional shutdown command */
	app.ShutdownCMD = "__shutdown"

	go func() {
		<-app.Shutdown
		app.Session.Disconnect()
		log.Println("finished")
	}()

	app.HandleIncomingEvents()
}
