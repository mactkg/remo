package remo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Slack interface {
	SetStatus(status Status) error
	PostMessage(to string, message string) (*MessageResponse, error)
	GetPermalink(channel string, ts string) (*PermalinkResponse, error)
}

// slack have informations to communicate with slack
type slack struct {
	baseURL string
	token   string
	client  http.Client
}

// Request struct for users.profile
type profileRequest struct {
	Profile Profile `json:"profile"`
}

// Request struct for chat.postMessage
type messageRequest struct {
	*Message
}

type permalinkRequest struct {
	Channel string
	Ts      string `json:"message_ts"`
}

// base struct for response from slack API
type response struct {
	Ok    bool
	Error string
}

type MessageResponse struct {
	response
	Channel string
	Ts      string
	Message Message
}

type PermalinkResponse struct {
	response
	Channel   string
	Permalink string
}

// Profile is representation of profile.
// It have only a embedded struct for now, but it could have some fields such as user's name or so.
type Profile struct {
	*Status
}

// Status is representation of user's status
type Status struct {
	Text       string `json:"status_text"`
	Emoji      string `json:"status_emoji"`
	Expiration int    `json:"status_expiration"`
}

// Message is representation of message on slack
type Message struct {
	Channel string `json:"channel"`
	Text    string `json:"text"`
	AsUser  bool   `json:"as_user"`
	Parse   string `json:"parse"`
}

var baseURL = "https://slack.com/api"

// NewSlack create instance of slack
func NewSlack(token string) Slack {
	client := http.Client{}
	return slack{
		baseURL: baseURL,
		token:   token,
		client:  client,
	}
}

func (s slack) createJsonRequest(method string, action string, data interface{}) (*http.Request, error) {
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

func (s slack) createURLEncodedRequest(method string, action string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, s.baseURL+"/"+action, body)
	if err != nil {
		return nil, err
	}

	if method == http.MethodPost {
		req.Header.Set("content-type", "application/x-www-form-urlencoded")
	}
	req.Header.Set("Authorization", "Bearer "+s.token)
	return req, nil
}

// SetStatus change status on slack
// required scope: users.profile:write
func (s slack) SetStatus(status Status) error {
	profile := Profile{&status}
	data := profileRequest{
		Profile: profile,
	}
	req, err := s.createJsonRequest("POST", "users.profile.set", data)
	if err != nil {
		return err
	}

	_, err = s.client.Do(req)
	if err != nil {
		return err
	}

	return nil
}

// PostMessage post a message to slack
// required scope: chat:write:user, chat:write:bot
// TODO: examinate what scopes are required specificaly[]
func (s slack) PostMessage(to string, message string) (*MessageResponse, error) {
	msg := Message{
		Text:    message,
		Channel: to,
		AsUser:  true,
		Parse:   "full",
	}
	data := messageRequest{&msg}
	req, err := s.createJsonRequest("POST", "chat.postMessage", data)
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	resData := MessageResponse{}
	resData.Ok = true
	err = json.NewDecoder(res.Body).Decode(&resData)
	if err != nil {
		return nil, err
	}

	if !resData.Ok {
		fmt.Printf("Error %v", resData)
		return nil, fmt.Errorf("error from slack API: %v", resData.Error)
	}

	return &resData, nil
}

func (s slack) GetPermalink(channel string, ts string) (*PermalinkResponse, error) {
	// This method does not currently accept application/json.
	query := fmt.Sprintf("?channel=%s&message_ts=%s", channel, ts)
	req, err := s.createURLEncodedRequest("GET", "chat.getPermalink"+query, &bytes.Buffer{})
	if err != nil {
		return nil, err
	}

	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}

	resData := &PermalinkResponse{}
	resData.Ok = true
	err = json.NewDecoder(res.Body).Decode(resData)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if !resData.Ok {
		return nil, fmt.Errorf("error: %v", resData.Error)

	}

	return resData, nil
}
