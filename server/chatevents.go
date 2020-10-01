package server

import (
	"../userdata"
	"strings"
)

func messageNewHandlerChat(user *userdata.User, chat *userdata.Chat, message map[string]interface{}) {
	var textObj = message["text"]
	var text = textObj.(string)
	if textObj == nil || len(text) < 1 {
		return
	}

	var supercall = text[0] == '/'
	if supercall {
		text = text[1:]
	}
	var parts = strings.Split(text, " ")
	var call = strings.ToLower(parts[0])

	if strings.Compare(call, "яркабот") == 0 {
		if len(parts) == 1 {
			userdata.SendChatMessage(chat, "Да да? Список команд: яркабот помощь")
		} else {
			if commands[parts[1]] != nil {
				if commands[parts[1]].Chat != nil {
					commands[parts[1]].Chat(user, chat, strings.Join(parts[1:], " "))
				} else {
					userdata.SendChatMessage(chat, "Эта команда в беседах не работает (")
				}
			} else {
				userdata.SendChatMessage(chat, "Список команд: яркабот помощь")
			}
		}
	} else if supercall {
		if commands[parts[0]] != nil {
			if commands[parts[0]].Chat != nil {
				commands[parts[0]].Chat(user, chat, text)
			} else {
				userdata.SendChatMessage(chat, "Эта команда в беседах не работает (")
			}
		} else {
			userdata.SendChatMessage(chat, "Список команд: яркабот помощь")
		}
	}
}