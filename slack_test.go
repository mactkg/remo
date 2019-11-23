package remo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSetStatus(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			reqProfile := profileRequest{}
			err := decoder.Decode(&reqProfile)
			if err != nil {
				t.Fatalf("Decode request body error: %v", err)
			}

			profile := reqProfile.Profile
			ok := true
			if !ok {
				t.Fatalf("Type assertion error: %v", reqProfile)
			}

			if r.URL.Path != "/users.profile.set" {
				t.Errorf("Wrong pass: %v", r.URL.Path)
			}

			if profile.Text != "Working" {
				t.Errorf("Wrong status text: %v", profile.Text)
			}

			if profile.Emoji != ":computer:" {
				t.Errorf("Wrong status text: %v", profile.Emoji)
			}

			w.Header().Set("content-Type", "text")
			fmt.Fprintf(w, "ok")
			return
		},
	))
	defer ts.Close()
	baseURL = ts.URL

	status := Status{Text: "Working", Emoji: ":computer:"}
	slack := NewSlack("token")
	err := slack.SetStatus(status)
	if err != nil {
		t.Fatalf("Request should be done: %v", err)
	}
}

func TestPostMessage(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			decoder := json.NewDecoder(r.Body)
			reqMessage := messageRequest{}
			err := decoder.Decode(&reqMessage)
			if err != nil {
				t.Fatalf("Decode request body error: %v", err)
			}

			if r.URL.Path != "/chat.postMessage" {
				t.Errorf("Wrong pass: %v", r.URL.Path)
			}

			if reqMessage.Text != "Hello!" {
				t.Errorf("Wrong text: %v", reqMessage.Text)
			}

			if reqMessage.Channel != "#random" {
				t.Errorf("Wrong post channel: %v", reqMessage.Channel)
			}

			if reqMessage.AsUser != true {
				t.Errorf("as_user must be true")
			}

			w.Header().Set("content-Type", "text")
			fmt.Fprintf(w, "ok")
			return
		},
	))
	defer ts.Close()
	baseURL = ts.URL

	slack := NewSlack("token")
	err := slack.PostMessage("#random", "Hello!")
	if err != nil {
		t.Fatalf("Request should be done: %v", err)
	}
}
