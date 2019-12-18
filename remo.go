package remo

import (
	"log"
	"os"
)

type Remo struct {
	slack  Slack
	config Config
}

var out = os.Stdout

func New(config *Config) *Remo {
	slack := NewSlack(config.Slack.Token)

	log.SetOutput(out)

	return &Remo{
		slack:  slack,
		config: *config,
	}
}

func (r *Remo) notice(postText, statusText, statusEmoji string) (err error) {
	err = r.slack.SetStatus(Status{
		Text:  statusText,
		Emoji: statusEmoji,
	})
	if err != nil {
		return
	}
	log.Printf("Success changing status to %s(%s)", statusText, statusEmoji)

	mainCh := r.config.Slack.MainPostChannel
	messageResponse, err := r.slack.PostMessage(mainCh, postText)
	if err != nil {
		return
	}
	log.Printf(`Success post "%s"(%s) to %s(%s)\n`, postText, messageResponse.Channel, mainCh, messageResponse.Ts)

	permalinkResponse, err := r.slack.GetPermalink(messageResponse.Channel, messageResponse.Ts)
	if err != nil {
		return err
	}
	log.Printf(`%s\n`, permalinkResponse.Permalink)

	for _, ch := range r.config.Slack.SubPostChannel {
		_, err := r.slack.PostMessage(ch, permalinkResponse.Permalink)
		if err != nil {
			log.Println(err)
		}
	}
	return
}

func (r *Remo) StartRemoteWork() error {
	return r.notice("リモートワーク開始します", "リモート中", ":male-technologist:")
}

func (r *Remo) PauseRemoteWork() error {
	return r.notice("休憩します", "離席中", ":tea:")
}

func (r *Remo) ResumeRemoteWork() error {
	return r.notice("戻りました", "リモート中", ":male-technologist:")
}

func (r *Remo) MoveToOffice() error {
	return r.notice("リモートワーク終了して出社します", "移動中", ":train:")
}

func (r *Remo) ArriveAtOffice() error {
	return r.notice("会社つきました", "", "")
}

func (r *Remo) FinishRemoteWork() error {
	return r.notice("リモートワーク終了します", "閉店", ":crescent_moon:")
}
