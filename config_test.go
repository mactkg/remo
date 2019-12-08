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
`
	buf := bytes.NewBufferString(data)

	config := Config{}
	dec := NewDecoder(buf)
	err := dec.Decode(&config)
	if err != nil {
		t.Fatal("Decoding sample data failed")
	}

	if config.Slack.Token != "foo" {
		t.Fatalf("wrong config.Slack.Token: got: %v, expected: foo", config.Slack.Token)
	}

	if config.Slack.MainPostChannel != "#kintai" {
		t.Fatalf("wrong config.Slack.MainPostChannel: got: %v, expected: #kintai", config.Slack.MainPostChannel)
	}

	sbcExpected := []string{"#times_mactkg", "#myproj_ds"}
	if !reflect.DeepEqual(config.Slack.SubPostChannel, sbcExpected) {
		t.Fatalf("wrong config.Slack.Token: got: %v, expected: %v", config.Slack.SubPostChannel, sbcExpected)
	}
}
