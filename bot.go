package main

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/cixtor/slackapi"
)

// Slackbot defines the application metadata.
type Slackbot struct {
	RobotID    string
	ChannelID  string
	Client     *slackapi.SlackAPI
	ReadmeFile string
	Startup    int
}

// NewSlackbot creates a new instance of the application.
func NewSlackbot() *Slackbot {
	startup := int(time.Now().Unix())
	log.Println("Starting at", startup)
	client := slackapi.New()

	binary, err := os.Executable()
	if err != nil {
		log.Fatal("help;\x20", err)
	}

	return &Slackbot{
		Client:     client,
		Startup:    startup,
		ReadmeFile: path.Dir(binary) + "/README.md",
	}
}

// HandleIncomingEvents processes the websocket events.
func (s *Slackbot) HandleIncomingEvents() {
	log.Println("Connecting to websocket")
	rtm, err := s.Client.NewRTM()

	if err != nil {
		log.Fatalf("RTM connection; %s", err.Error())
		return
	}

	log.Println("Listening to RTM events")
	go rtm.ManageEvents()

	for msg := range rtm.Events {
		switch event := msg.Data.(type) {
		case *slackapi.MessageEvent:
			s.HandleMessage(event)
		}
	}
}
