package main

import (
	"fmt"

	"github.com/tinti/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("mytoken")
	if err != nil {
		panic(err)
	}

	url := "amqp://guest:guest@localhost:5672/guest"
	go tgbotapi.SimpleServer(url, bot)
	r, err := tgbotapi.RemoteBotDial(url)
	if err != nil {
		panic(err)
	}
	defer tgbotapi.RemoteBotClose(r)

	me, err := r.GetMe()
	fmt.Println(me, err)
}
