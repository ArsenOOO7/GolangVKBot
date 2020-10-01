package vkhttp

import (
	"../util"
	"fmt"
	"io/ioutil"
	"net/http"
)

type VKResponse struct {
	Response []byte
}

type VKRequest struct {
	Method string
	Params []VKParam
}

type VKParam struct {
	Key string
	Value string
}

func ParseConfig(config *util.Config) ([]VKParam) {
	return []VKParam{
		{
			Key:   "access_token",
			Value: config.AccessToken,
		},
		{
			Key:   "v",
			Value: config.Version,
		},
	}
}

func SendRequest(request VKRequest) *VKResponse  {
	response, err := http.Get("https://api.vk.com/method/" + request.Method + buildHttpParams(request.Params))
	if err != nil {
		fmt.Print(err)
		return nil
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return &VKResponse{Response:body}
}

func SendRawRequest(request VKRequest) *VKResponse {
	response, err := http.Get(request.Method + buildHttpParams(request.Params))
	if err != nil {
		fmt.Print(err)
		return nil
	}

	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)

	return &VKResponse{Response:body}
}

func buildHttpParams(params []VKParam) string {
	var result = "?"

	for _, param := range params {
		result += param.Key + "=" + param.Value + "&"
	}

	return result[:len(result) - 1]
}
