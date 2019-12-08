package remo

import (
	"github.com/BurntSushi/toml"
	"io"
)

type Config struct {
	Slack SlackConfig
}

type SlackConfig struct {
	Token           string
	MainPostChannel string
	SubPostChannel  []string
}

type Decoder struct {
	in io.Reader
}

func NewDecoder(in io.Reader) *Decoder {
	return &Decoder{in}
}

func (d Decoder)Decode(v *Config) (err error) {
	_, err = toml.DecodeReader(d.in, v)
	return
}
