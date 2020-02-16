package remo

import (
	"bytes"
	"reflect"
	"testing"
)

func TestDecoder(t *testing.T) {
	data := `
[slack]
token = "foo"
mainPostChannel = "#kintai"
subPostChannel = [
  "#times_mactkg",
  "#myproj_ds"
]

[remote]
	[remote.start]
	post = "リモートワーク開始します"
	state_emoji = ":male-technologist:"
	state_message = "リモート中"

	[remote.leave]
	post = "離席します"
	state_emoji = ":tea:"
	state_message = "離席中"

	[remote.back]
	post = "戻りました"
	state_emoji = ":male-technologist:"
	state_message = "リモート中"

	[remote.move]
	post = "リモートワーク終了して出社します"
	state_emoji = ":train:"
	state_message = "移動中"
	
	[remote.end]
	post = "リモートワーク終了します"
	state_emoji = ":crescent_moon:"
	state_message = "閉店"

[office]
	[office.start]
	post = "出社しました"
	state_emoji = ":male-technologist:"
	state_message = "勤務中"

	[office.leave]
	# コメントアウトしてある場合は、基本remoteが引き継がれる
	# post = "離席します"
	# state_emoji = ":tea:"
	# state_message = "離席中"

	[office.back]
	post = "戻りました"
	state_emoji = ":male-technologist:"
	state_message = "勤務中"

	[office.move]
	post = "移動します"
	state_emoji = ":train:"
	state_message = "移動中"
	
	[office.end]
	post = "帰ります"
	state_emoji = ":crescent_moon:"
	state_message = "閉店"
`
	buf := bytes.NewBufferString(data)

	config := Config{}
	dec := NewDecoder(buf)
	err := dec.Decode(&config)
	if err != nil {
		t.Fatal("Decoding sample data failed")
	}

	if config.Slack.Token != "foo" {
		t.Fatalf("wrong config.slack.Token: got: %v, expected: foo", config.Slack.Token)
	}

	if config.Slack.MainPostChannel != "#kintai" {
		t.Fatalf("wrong config.slack.MainPostChannel: got: %v, expected: #kintai", config.Slack.MainPostChannel)
	}

	sbcExpected := []string{"#times_mactkg", "#myproj_ds"}
	if !reflect.DeepEqual(config.Slack.SubPostChannel, sbcExpected) {
		t.Fatalf("wrong config.slack.Token: got: %v, expected: %v", config.Slack.SubPostChannel, sbcExpected)
	}
}
