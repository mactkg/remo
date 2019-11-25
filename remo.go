package remo

type Remo struct {
	slack  Slack
	Config Config
}

type Config struct {
	SlackToken      string
	MainPostChannel string
}

func New(config Config) *Remo {
	slack := NewSlack(config.SlackToken)
	return &Remo{
		slack:  slack,
		Config: config,
	}
}

func (r *Remo) StartRemoteWork() {
	r.slack.SetStatus(Status{
		Text:  "リモート中",
		Emoji: ":male-technologist:",
	})
	r.slack.PostMessage(r.Config.MainPostChannel, "リモートワーク開始します")
}

func (r *Remo) PauseRemoteWork() {
	r.slack.SetStatus(Status{
		Text:  "離席中",
		Emoji: ":tea:",
	})
	r.slack.PostMessage(r.Config.MainPostChannel, "休憩します")
}

func (r *Remo) ResumeRemoteWork() {
	r.slack.SetStatus(Status{
		Text:  "リモート中",
		Emoji: ":male-technologist:",
	})
	r.slack.PostMessage(r.Config.MainPostChannel, "戻りました")
}

func (r *Remo) MoveToOffice() {
	r.slack.SetStatus(Status{
		Text:  "移動中",
		Emoji: ":train:",
	})
	r.slack.PostMessage(r.Config.MainPostChannel, "リモートワーク終了して出社します")
}

func (r *Remo) FinishRemoteWork() {
	r.slack.SetStatus(Status{
		Text:  "閉店",
		Emoji: ":crescent_moon:",
	})
	r.slack.PostMessage(r.Config.MainPostChannel, "リモートワーク終了します")
}
