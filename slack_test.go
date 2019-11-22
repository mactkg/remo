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
			status := Status{}
			err := decoder.Decode(&status)
			if err != nil {
				t.Fatalf("Decode request body error: %v", err)
			}

			if r.URL.Path != "/users.profile.set" {
				t.Errorf("Wrong pass: %v", r.URL.Path)
			}

			if status.Text != "Working" {
				t.Errorf("Wrong status text: %v", status.Text)
			}

			if status.Emoji != ":computer:" {
				t.Errorf("Wrong status text: %v", status.Emoji)
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
			message := Message{}
			err := decoder.Decode(&message)
			if err != nil {
				t.Fatalf("Decode request body error: %v", err)
			}

			if r.URL.Path != "/chat.postMessage" {
				t.Errorf("Wrong pass: %v", r.URL.Path)
			}

			if message.Text != "Hello!" {
				t.Errorf("Wrong text: %v", message.Text)
			}

			if message.Channel != "#random" {
				t.Errorf("Wrong post channel: %v", message.Channel)
			}

			if message.AsUser != true {
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
