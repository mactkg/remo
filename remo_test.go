package remo

import "testing"

type SlackMock struct {
	status Status
	to	string
	message string
	channel string
	ts string
}

func (s SlackMock) SetStatus(status Status) error {
	s.status = status
	return nil
}

func (s SlackMock) PostMessage(to string, message string) (*MessageResponse, error) {
	s.to = to
	s.message = message
	return nil, nil
}

func (s SlackMock) GetPermalink(channel string, ts string) (*PermalinkResponse, error) {
	s.channel = channel
	s.ts = ts
	return nil, nil
}

var config Config
var slackMock SlackMock
var remo Remo

func init() {
	config = Config{
		Slack: SlackConfig{
			Token:      "dummy-token",
			MainPostChannel: "#random",
		},
	}
	slackMock = SlackMock{}
	remo = Remo{
		slack: slackMock,
		config: config,
	}
}

func TestNew(t *testing.T) {
	remo := New(&config)
	if remo == nil {
		t.Fatal("remo.New() should be returned pointer of Remo{}")
	}
}

func TestStart(t *testing.T) {
	remo.StartRemoteWork()
	if slackMock.channel != config.Slack.MainPostChannel {
		t.Fatal("Should post to " + config.Slack.MainPostChannel)
	}

}
