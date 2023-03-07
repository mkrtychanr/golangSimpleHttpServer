package server

import (
	"encoding/json"
	"os"
)

type config struct {
	Host             string `json:"host"`
	Port             string `json:"port"`
	SubscribeSubject string `json:"subscribe_subject"`
}

type configWrap struct {
	Config config `json:"server"`
}

func newConfig(path string) (*config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := configWrap{}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}
	return &config.Config, nil
}
