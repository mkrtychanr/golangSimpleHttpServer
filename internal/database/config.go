package database

import (
	"encoding/json"
	"os"
)

type config struct {
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func newConfig(path string) (*config, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	config := config{}
	err = json.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
