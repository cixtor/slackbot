package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cixtor/slackapi"
)

// HandleMessage reacts a message from the monitored channel.
func (s *Slackbot) HandleMessage(event *slackapi.MessageEvent) {
	/* detect bot replies */
	if event.User == s.RobotID {
		s.LogMessage("bot", event)
		return
	}

	/* ignore deleted messages; empty text */
	if event.Subtype == "message_deleted" {
		return
	}

	/* detect messages and commands sent before the service started */
	timestamp := event.Timestamp[0:strings.Index(event.Timestamp, ".")]
	if number, err := strconv.Atoi(timestamp); err != nil || number < s.Startup {
		log.Printf(
			"msg; [old] %s %s: %s\n",
			event.Timestamp,
			event.User,
			event.Text)
		return
	}

	if s.HandleHelp(event) {
		s.LogMessage("cmd", event)
		return
	}

	if s.HandleUptime(event) {
		s.LogMessage("cmd", event)
		return
	}

	s.LogMessage("new", event)
}

// HandleHelp reacts to an uptime request.
func (s *Slackbot) HandleHelp(event *slackapi.MessageEvent) bool {
	if event.Text == "help" {
		file, err := os.Open(s.ReadmeFile)
		if err != nil {
			log.Println("readme open;\x20", err)
			return false
		}

		body, err := ioutil.ReadAll(file)
		if err != nil {
			log.Println("readme read;\x20", err)
			return false
		}

		reply := string(body)
		reHeading := regexp.MustCompile(`### (.+)`)
		reMention := regexp.MustCompile(`\*\*([@#]\S+)\*\*`)
		reAnchors := regexp.MustCompile(`\[([^\]]+)\]\(([^\)]+)\)`)

		reply = reHeading.ReplaceAllString(reply, `*$1*`)
		reply = reMention.ReplaceAllString(reply, `$1`)
		reply = reAnchors.ReplaceAllString(reply, `<$2|$1>`)

		session := s.Client.InstantMessageOpen(event.User)
		s.Client.ChatPostMessage(session.Channel.ID, reply)
		return true
	}

	return false
}

// HandleUptime reacts to an uptime request.
func (s *Slackbot) HandleUptime(event *slackapi.MessageEvent) bool {
	if event.Text == "uptime" {
		uptime := time.Since(time.Unix(int64(s.Startup), 0))
		reply := fmt.Sprintf("Running since %s", uptime)
		s.Client.ChatPostMessage(event.Channel, reply)
		return true
	}

	return false
}
