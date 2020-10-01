package userdata

import (
	"../util"
	"../vkhttp"
	"net/url"
	"strconv"
)

func SendChatMessage(chat *Chat, text string) (*vkhttp.VKResponse)  {
	request := vkhttp.VKRequest{
		Method: "messages.send",
		Params: append(vkhttp.ParseConfig(&util.DefaultConfig), vkhttp.VKParam{
			Key:   "chat_id",
			Value: strconv.Itoa(int(chat.Id)),
		}, vkhttp.VKParam{
			Key:   "message",
			Value: url.QueryEscape(text),
		}, vkhttp.VKParam{
			Key:   "random_id",
			Value: "0",
		}),
	}

	return vkhttp.SendRequest(request)
}

func SendChatMessageWithAttachments(chat *Chat, text string, attachments []*Attachment) (*vkhttp.VKResponse)  {
	request := vkhttp.VKRequest{
		Method: "messages.send",
		Params: append(vkhttp.ParseConfig(&util.DefaultConfig), vkhttp.VKParam{
			Key:   "chat_id",
			Value: strconv.Itoa(int(chat.Id)),
		}, vkhttp.VKParam{
			Key:   "message",
			Value: url.QueryEscape(text),
		}, vkhttp.VKParam{
			Key:   "random_id",
			Value: "0",
		}, vkhttp.VKParam{
			Key:   "attachment",
			Value: encodeAttachments(attachments),
		}),
	}

	return vkhttp.SendRequest(request)
}