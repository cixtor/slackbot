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

// LogRobotMessage writes the requested command to the system logs.
func (s *Slackbot) LogRobotMessage(event *slackapi.MessageEvent) {
	log.Printf(
		"msg; [rep] %s at %s\n",
		event.Timestamp,
		event.Channel)
}

// LogCommand writes the requested command to the system logs.
func (s *Slackbot) LogCommand(event *slackapi.MessageEvent) {
	log.Printf(
		"msg; [cmd] %s %s: %s\n",
		event.Timestamp,
		event.User,
		event.Text)
}

// LogMessage writes the user message to the system logs.
func (s *Slackbot) LogMessage(event *slackapi.MessageEvent) {
	/* monitor other messages (cut to 76 chars) */
	if len(event.Text) > 76 {
		log.Printf(
			"msg; [new] %s %s: %s...\n",
			event.Timestamp,
			event.User,
			event.Text[0:76])
		return
	}

	/* monitor other messages (with less than 76 characters) */
	log.Printf(
		"msg; [new] %s %s: %s\n",
		event.Timestamp,
		event.User,
		event.Text)
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
