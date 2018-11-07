package messenger

import (
	"fmt"

	"github.com/nlopes/slack"
)

// NewSlackMessenger returns a new Slack messenger
func NewSlackMessenger(channel, message, apikey string) (*SlackMessenger, error) {
	if apikey == "" {
		return nil, fmt.Errorf("no api key provided")
	}
	if channel == "" {
		return nil, fmt.Errorf("no channel (chat_id) provided")
	}
	sm := new(SlackMessenger)
	sm.Message = message
	sm.Channel = channel
	sm.client = slack.New(apikey)

	return sm, nil
}

// SlackMessenger represents a messenger for Slack
type SlackMessenger struct {
	Channel string
	Message string
	client  *slack.Client
}

// SendMessage implements messenger.SendMessage
func (s *SlackMessenger) SendMessage() error {
	fmt.Println("sending Slack message...")
	channelID, timestamp, err := s.client.PostMessage(
		s.Channel,
		s.Message,
		slack.PostMessageParameters{})

	if err != nil {
		err = fmt.Errorf("Slack messenger: Something went wrong sending a message (%s at %s): %s", channelID, timestamp, err)
	}

	return err
}