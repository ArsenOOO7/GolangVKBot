package server

import (
	"../threadpool"
	"../util"
	"../vkhttp"
	"encoding/json"
	"fmt"
	"strconv"
)

type LongPollServer struct {
	Server string `json:"server"`
	Ts int `json:"ts"`
	Key string `json:"key"`
}

func Start(config *util.Config)  {
	var attempts = 1
	start:
	serverData, err := initServer(config)
	if err != nil {
		fmt.Println("Init error: ", err, "| Attempt", attempts, "of 5")
		if attempts == 5 {
			fmt.Println("Disabling bot. Too much bad initializations")
			return
		}

		attempts++
		goto start
	}

	type Updates struct {
		Ts string `json:"ts"`
		Updates []map[string]string `json:"updates"`
	}

	for true {
		request := vkhttp.VKRequest{
			Method: serverData.Server,
			Params: append(vkhttp.ParseConfig(config), vkhttp.VKParam{
				Key:   "act",
				Value: "a_check",
			}, vkhttp.VKParam{
				Key:   "key",
				Value: serverData.Key,
			}, vkhttp.VKParam{
				Key:   "wait",
				Value: "10",
			}, vkhttp.VKParam{
				Key:   "ts",
				Value: strconv.Itoa(serverData.Ts),
			}),
		}

		var newTs ts
		response := vkhttp.SendRawRequest(request)
		err := json.Unmarshal(response.Response, &newTs)
		if err != nil {
			fmt.Println(err)
			continue
		}

		serverData.Ts, err = strconv.Atoi(newTs.Ts)
		if err != nil {
			fmt.Println(err)
			continue
		}

		var updates updates
		err = json.Unmarshal(response.Response, &updates)
		if err != nil {
			fmt.Println(err)
			continue
		}

		for _, update := range updates.Updates {
			mapped, ok := update.(map[string]interface{})
			if !ok || mapped == nil {
				continue
			}

			threadpool.Tasks = append(threadpool.Tasks, &threadpool.Task{Task:mapped})
			threadpool.SigWork = true
		}
	}
}

func initServer(config *util.Config) (*LongPollServer, error) {
	request := vkhttp.VKRequest{
		Method: "groups.getLongPollServer",
		Params: append(vkhttp.ParseConfig(config), vkhttp.VKParam{
			Key:   "group_id",
			Value: fmt.Sprintf("%v", config.GroupId),
		}),
	}

	response := vkhttp.SendRequest(request)
	serverData := map[string]map[string]string{}

	err := json.Unmarshal(response.Response, &serverData)
	if err != nil {
		return nil, err
	}

	if _, ok := serverData["response"]; !ok {
		return nil, fmt.Errorf("Cannot find a correct response")
	}

	convertedTs, err := strconv.Atoi(serverData["response"]["ts"])
	if err != nil {
		return nil, err
	}

	fmt.Println("Server initialized!")

	return &LongPollServer{
		Server: serverData["response"]["server"],
		Ts:     convertedTs,
		Key:    serverData["response"]["key"],
	}, nil
}