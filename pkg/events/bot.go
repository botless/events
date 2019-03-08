package events

import (
	"fmt"
	"log"
	"strings"
)

const (
	bot_type_template = "botless.bot.%s" // type
)

var knownBotEvents = []string{
	"response",
}

type Bot int

// Message, start simple.
type Message struct {
	Channel string `json:"channel,omitempty"`
	Text    string `json:"text,omitempty"`
}

func (Bot) Type(t string) string {
	if contains(knownBotEvents, t) {
		log.Printf("[WARN] unknown bot event type: %q", t)
	}
	return strings.ToLower(fmt.Sprintf(bot_type_template, t))
}
