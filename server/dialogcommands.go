package server

import (
	"../threadpool"
	"../userdata"
	"../util"
	"fmt"
	"runtime"
	"strconv"
)

func handleStatDialog(user *userdata.User, command string)  {
	text := "Количество свободных воркеров: " + strconv.Itoa(threadpool.CountFreeWorkers()) + "/" + strconv.Itoa(threadpool.WorkerCount) + "\n"

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	text += fmt.Sprintf("Мной сожрано: %v KiB", bToKiB(m.TotalAlloc)) + fmt.Sprintf("\nМне разрешили сожрать: %v KiB", bToKiB(m.Sys))

	userdata.SendMessage(user, text)
}

func handlePukDialog(user *userdata.User, command string)  {
	userdata.SendMessage(user, "Ты сделал пук и в ответ тебе в лицо прилетел смачный ВЫПУК")
}

func handleDopolniDialog(user *userdata.User, command string)  {

}

func handleHelpDialog(user *userdata.User, command string)  {
	text := "Команды бота:\n"
	for cmd, info := range help {
		text += "\n    " + cmd + " - " + info
	}
	text += "\n\nМеня можно называть: яркабот\nАвтор бота - [id572440179|Yarka Ilin]"

	userdata.SendMessage(user, text)
}

func handleAboutDialog(user *userdata.User, command string)  {
	userdata.SendMessage(user, "JustYarka Bot\n\nНаписан на языке: GoLang\nБиблиотеки: Библиотеки GoLang\nВерсия VK API: " + util.DefaultConfig.Version + "\nВерсия бота: " + util.DefaultConfig.BotVersion)

}
