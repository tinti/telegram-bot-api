package main

import (
	"fmt"
	"net/url"

	"github.com/tinti/telegram-bot-api"
)

func main() {
	go tgbotapi.DummyServer()
	r := tgbotapi.RemoteBotAPI{}

	endpoint := "hi"
	params := url.Values{}

	apiResponse, err := r.MakeRequest(endpoint, params)

	fmt.Println(apiResponse, err)
}
