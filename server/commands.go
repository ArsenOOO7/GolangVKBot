package server

var commands = map[string]*CommandHandler{
	"стат": {
		Chat:	handleStatChat,
		Dialog:	handleStatDialog,
	},
	"пук": {
		Chat:	handlePukChat,
		Dialog: handlePukDialog,
	},
	"дополни": {
		Chat: 	handleDopolniChat,
		Dialog: handleDopolniDialog,
	},
	"помощь": {
		Chat: 	handleHelpChat,
		Dialog: handleHelpDialog,
	},
	"about": {
		Chat:	handleAboutChat,
		Dialog:	handleAboutDialog,
	},
}

var help = map[string]string {
	"стат": "Статистика бота",
	"пук": "Пукнуть и получить смачный ответ",
	"дополни": "Дополнить свой текст",
	"помощь": "Из названия команды думаю догадаешься ¯\\_(ツ)_/¯",
	"about": "Инфа о боте",
}

