package server

import (
	"../userdata"
)

var handlers = map[string]*EventHandler {
	"message_new": {
		Chat:   messageNewHandlerChat,
		Dialog: messageNewHandlerDialog,
	},
}

func Handle(update map[string]interface{})  {
	var objectInterface = update["object"]
	if objectInterface != nil {
		objectMap, ok := objectInterface.(map[string]interface{})
		if !ok || objectMap == nil {
			return
		}

		var eventType = update["type"]
		var eventTypeString = eventType.(string)
		if eventType != nil && handlers[eventTypeString] != nil {
			mappedMessage, ok := objectMap["message"].(map[string]interface{})
			if !ok || mappedMessage == nil {
				return
			}

			var id = mappedMessage["from_id"]
			var peerId = int64(mappedMessage["peer_id"].(float64))

			if peerId > 2000000000 {
				handlers[eventTypeString].Chat(&userdata.User{
					Id: int64(id.(float64)),
					IsUser: int64(id.(float64)) > 0,
				}, &userdata.Chat{Id:peerId - 2000000000}, mappedMessage)


			} else {
				handlers[eventTypeString].Dialog(&userdata.User{
					Id: int64(id.(float64)),
					IsUser: int64(id.(float64)) > 0,
				}, mappedMessage)
			}
		}
	}
}




