package server

import (
	"../userdata"
)

type updates struct {
	Updates []interface{} `json:"updates"`
}

type ts struct {
	Ts string `json:"ts"`
}

type CommandHandler struct {
	Chat func(*userdata.User, *userdata.Chat, string)
	Dialog func(*userdata.User, string)
}

type EventHandler struct {
	Chat func(*userdata.User, *userdata.Chat, map[string]interface{})
	Dialog func(*userdata.User, map[string]interface{})
}


