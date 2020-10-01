package util

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	AccessToken string `json:"access_token"`
	GroupId     int64  `json:"group_id"`
	Version     string `json:"version"`
	GoroutineCount int `json:"default_goroutine_count"`
	BotVersion  string `json:"bot_version"`
}

var DefaultConfig Config

func CreateConfig() (error)  {
	target := Config{}

	bytes, err := ioutil.ReadFile("resources" + string(filepath.Separator) + "config.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &target)
	if err != nil {
		return err
	}

	DefaultConfig = target
	return nil
}
