package main

import (
	"log"
	"os"
)

func main() {
	app := NewSlackbot()

	log.Println("Setting session token")
	app.Client.SetToken(os.Getenv("SLACK_TOKEN"))

	log.Println("Identifying robot ID")
	resp := app.Client.AuthTest()
	app.RobotID = resp.UserID

	app.HandleIncomingEvents()
}
