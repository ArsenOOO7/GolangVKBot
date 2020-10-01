package server

import (
	"../threadpool"
	"../userdata"
	"../util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strconv"
	"strings"
)

func handleStatChat(user *userdata.User, chat *userdata.Chat, command string)  {
	text := "Количество свободных воркеров: " + strconv.Itoa(threadpool.CountFreeWorkers()) + "/" + strconv.Itoa(threadpool.WorkerCount) + "\n"

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	text += fmt.Sprintf("Мной сожрано: %v KiB", bToKiB(m.TotalAlloc)) + fmt.Sprintf("\nМне разрешили сожрать: %v KiB", bToKiB(m.Sys))

	userdata.SendChatMessage(chat, text)
}

func handlePukChat(user *userdata.User, chat *userdata.Chat, command string)  {
	userdata.SendChatMessage(chat, "Ты сделал пук и в ответ тебе в лицо прилетел смачный ВЫПУК")
}

func handleDopolniChat(user *userdata.User, chat *userdata.Chat, command string)  {
	values := map[string]interface{}{"num_samples": 1, "length": 30, "prompt": strings.Join(strings.Split(command, " ")[1:], "")}
	fmt.Println(values["prompt"])
	jsonData, _ := json.Marshal(values)
	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(jsonData)

	response, err := http.Post("https://models.dobro.ai/gpt2/medium/", "application/json", buf)
	if err != nil {
		userdata.SendChatMessage(chat, "Ой какая то ашипка (")
	} else {
		defer response.Body.Close()

		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(body))
	}
}

func handleHelpChat(user *userdata.User, chat *userdata.Chat, command string) {
	text := "Команды бота:\n"
	for cmd, info := range help {
		text += "\n    " + cmd + " - " + info
	}
	text += "\n\nМеня можно называть: яркабот\nАвтор бота - vk.com/just_yarka (Yarka Ilin)"

	userdata.SendChatMessage(chat, text)
}

func handleAboutChat(user *userdata.User, chat *userdata.Chat, command string)  {
	userdata.SendChatMessage(chat, "JustYarka Bot\n\nНаписан на языке: GoLang\nБиблиотеки: Библиотеки GoLang\nВерсия VK API: " + util.DefaultConfig.Version + "\nВерсия бота: " + util.DefaultConfig.BotVersion)
}


func bToKiB(b uint64) uint64 {
	return b / 1024
}