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

	endpoint = "hey"
	params2 := map[string]string{}
	fieldname := ""
	file := interface{}(nil)

	apiResponse, err = r.UploadFile(endpoint, params2, fieldname, file)
	fmt.Println(apiResponse, err)
}
