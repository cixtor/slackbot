package main

import (
	"log"
	"strconv"
	"strings"

	"github.com/cixtor/slackapi"
)

// HandleMessage reacts a message from the monitored channel.
func (s *Slackbot) HandleMessage(event *slackapi.MessageEvent) {
	/* detect bot replies */
	if event.User == s.RobotID {
		s.LogRobotMessage(event)
		return
	}

	/* detect when a message is deleted */
	if event.Subtype == "message_deleted" {
		log.Printf("msg; [del] %s\n", event.Timestamp)
		return
	}

	/* detect messages and commands sent before the service started */
	timestamp := event.Timestamp[0:strings.Index(event.Timestamp, ".")]
	if number, err := strconv.Atoi(timestamp); err != nil || number < s.Startup {
		log.Printf("msg; [old] %s %s: %s\n", event.Timestamp, event.User, event.Text)
		return
	}

	s.LogMessage(event)
}
