package remo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// Slack have informations to communicate with Slack
type Slack struct {
	baseURL string
	token   string
	client  http.Client
}

type ProfileRequest struct {
	Profile Profile `json:"profile"`
}

type Profile struct {
	*Status
}

// Status is representation of user's status
type Status struct {
	Text       string `json:"status_text"`
	Emoji      string `json:"status_emoji"`
	Expiration int    `json:"status_expiration"`
}

// Message is representation of message on Slack
type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	AsUser  bool   `json:"as_user"`
	Parse   string `json:"parse"`
}

var baseURL = "https://slack.com/api"

// NewSlack create instance of Slack
func NewSlack(token string) Slack {
	client := http.Client{}
	return Slack{
		baseURL: baseURL,
		token:   token,
		client:  client,
	}
}

func (s Slack) createRequest(method string, action string, data interface{}) (*http.Request, error) {
	b, err := json.Marshal(&data)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	req, err := http.NewRequest(method, s.baseURL+"/"+action, buf)
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/json")
	req.Header.Set("Authorization", "Bearer "+s.token)
	return req, nil
}

// SetStatus change status on Slack
// required scope: users.profile:write
func (s Slack) SetStatus(status Status) error {
	profile := Profile{&status}
	data := ProfileRequest{
		Profile: profile,
	}
	req, err := s.createRequest("POST", "users.profile.set", data)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// PostMessage post a message to Slack
// required scope: chat:write:user, chat:write:bot
// TODO: examinate what scopes are required specificaly[]
func (s Slack) PostMessage(to string, message string) error {
	msg := Message{
		Text:    message,
		Channel: to,
		AsUser:  true,
		Parse:   "full",
	}
	req, err := s.createRequest("POST", "chat.postMessage", msg)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
