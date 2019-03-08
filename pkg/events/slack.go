package events

import (
	"fmt"
	"log"
	"strings"

	"github.com/cloudevents/sdk-go/pkg/cloudevents/types"
)

const (
	slack_source_channel_template = "https://%s.slack.com/messages/%s/" // domain, channel
	slack_source_domain_template  = "https://%s.slack.com/"             // domain
	slack_type_template           = "botless.slack.%s"                  // type
)

var slackKnownEvents = []string{
	"welcome",
	"message",
	"latency",
}

type Slack int

func (Slack) Type(t string) string {
	if contains(slackKnownEvents, t) {
		log.Printf("[WARN] unknown slack event type: %q", t)
	}
	return strings.ToLower(fmt.Sprintf(slack_type_template, t))
}

func (Slack) SourceForDomain(domain string) types.URLRef {
	source := types.ParseURLRef(fmt.Sprintf(slack_source_domain_template, domain))
	if source == nil {
		return types.URLRef{}
	}
	return *source
}

func (Slack) SourceForChannel(domain, channel string) types.URLRef {
	source := types.ParseURLRef(fmt.Sprintf(slack_source_channel_template, domain, channel))
	if source == nil {
		return types.URLRef{}
	}
	return *source
}
