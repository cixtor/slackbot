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
	Startup     int
	RobotID     string
	ChannelID   string
	ReadmeFile  string
	ShutdownCMD string
	Shutdown    chan bool
	Session     *slackapi.RTM
	Client      *slackapi.SlackAPI
}

// NewSlackbot creates a new instance of the application.
func NewSlackbot(token string, cookie string) *Slackbot {
	startup := int(time.Now().Unix())

	log.Println("Starting at", startup)

	client := slackapi.New()
	client.SetToken(token)
	client.SetCookie(cookie)

	log.Println("Identifying robot ID")

	owner, err := client.AuthTest()

	if err != nil {
		log.Fatal("auth test;\x20", err)
	}

	if !owner.Ok {
		log.Fatal("auth test;\x20", owner.Error)
	}

	binary, err := os.Executable()

	if err != nil {
		log.Fatal("help;\x20", err)
	}

	return &Slackbot{
		Client:     client,
		Startup:    startup,
		Shutdown:   make(chan bool, 1),
		RobotID:    owner.UserID,
		ReadmeFile: path.Dir(binary) + "/README.md",
	}
}

// HandleIncomingEvents processes the websocket events.
func (s *Slackbot) HandleIncomingEvents() {
	log.Println("Connecting to websocket")

	rtm, err := s.Client.NewRTM(slackapi.RTMInput{})

	if err != nil {
		log.Fatal("RTM connection;\x20", err)
	}

	s.Session = rtm

	log.Println("Listening to RTM events")

	s.Session.ManageEvents()

	for msg := range s.Session.Events {
		switch event := msg.Data.(type) {
		case *slackapi.MessageEvent:
			s.HandleMessage(event)
		default:
			log.Printf("%#v\n", msg)
		}
	}
}

// LogMessage writes the user message to the system logs.
func (s *Slackbot) LogMessage(tipo string, event *slackapi.MessageEvent) {
	// Monitor other messages (cut to 76 characters).
	if len(event.Text) > 76 {
		log.Printf(
			"msg; [%s] %s %s: %s...\n",
			tipo,
			event.Timestamp,
			event.User,
			event.Text[0:76],
		)
		return
	}

	// Monitor other messages (with less than 76 characters).
	log.Printf(
		"msg; [%s] %s %s: %s\n",
		tipo,
		event.Timestamp,
		event.User,
		event.Text,
	)
}
