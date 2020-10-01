package userdata

import (
	"../util"
	"../vkhttp"
	"net/url"
	"strconv"
)


func SendMessage(user *User, text string) (*vkhttp.VKResponse)  {
	request := vkhttp.VKRequest{
		Method: "messages.send",
		Params: append(vkhttp.ParseConfig(&util.DefaultConfig), vkhttp.VKParam{
			Key:   "user_id",
			Value: strconv.Itoa(int(user.Id)),
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

func SendMessageWithAttachments(user *User, text string, attachments []*Attachment) (*vkhttp.VKResponse)  {
	request := vkhttp.VKRequest{
		Method: "messages.send",
		Params: append(vkhttp.ParseConfig(&util.DefaultConfig), vkhttp.VKParam{
			Key:   "user_id",
			Value: strconv.Itoa(int(user.Id)),
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


