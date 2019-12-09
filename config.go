package remo

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"io"
	"os"
	"path"
)

var defaultConfig = `
[slack]
# token = "xoxp-xxx"
# mainPostChannel = '#report'
# subPostChannel [
#	"#my_personal_channel"
#]
`

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

func (d Decoder) Decode(v *Config) (err error) {
	_, err = toml.DecodeReader(d.in, v)
	return
}

func CreateConfigFile() (*os.File, error) {
	configPath, err := GetDefaultConfigPath()
	if err != nil {
		return nil, err
	}

	// check file already exists
	_, err = os.Open(configPath)
	if err == nil {
		return nil, fmt.Errorf("the default config file already exists")
	}

	configDir := path.Dir(configPath)
	err = os.MkdirAll(configDir, 0755)
	if err != nil {
		return nil, fmt.Errorf("can't create dirctory: %v", configDir)
	}

	flg := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	config, err := os.OpenFile(configPath, flg, 0600)
	if err != nil {
		return nil, fmt.Errorf("can't create file: %v, %v", configPath, err)
	}

	_, err = config.WriteString(defaultConfig)
	if err != nil {
		return config, fmt.Errorf("file created, but can't write default config data")
	}

	return config, nil
}

func GetDefaultConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	return homeDir + "/.config/remo/config.toml", nil
}
