package server

import (
	"../userdata"
	"strings"
)

func messageNewHandlerDialog(user *userdata.User, message map[string]interface{}) {
	var textObj = message["text"]
	var text = textObj.(string)
	if textObj == nil || len(text) < 1 {
		return
	}

	if text[0] == '/' {
		text = text[1:]
	}
	var parts = strings.Split(text, " ")
	var call = strings.ToLower(parts[0])

	if strings.Compare(call, "яркабот") == 0 {
		if len(parts) == 1 {
			userdata.SendMessage(user, "Да да? Список команд: помощь")
		} else {
			if commands[parts[1]] != nil {
				if commands[parts[1]].Dialog != nil {
					commands[parts[1]].Dialog(user, strings.Join(parts[1:], " "))
				} else {
					userdata.SendMessage(user, "Эта команда не работает в личных сообщениях (")
				}
			} else {
				userdata.SendMessage(user, "Список команд: помощь")
			}
		}
	} else {
		if commands[parts[0]] != nil {
			if commands[parts[0]].Dialog != nil {
				commands[parts[0]].Dialog(user, text)
			} else {
				userdata.SendMessage(user, "Эта команда не работает в личных сообщениях (")
			}
		} else {
			userdata.SendMessage(user, "Список команд: помощь")
		}
	}
}
