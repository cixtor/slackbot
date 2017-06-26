package main

import (
	"log"
	"os"
)

func main() {
	app := NewSlackbot(os.Getenv("SLACK_TOKEN"))

	app.HandleIncomingEvents()

	log.Println("finished")
}
