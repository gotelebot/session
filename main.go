package main

import (
	"time"

	tele "github.com/3JoB/telebot"
	"github.com/3JoB/telebot/middleware"
)

func main() {
	pref := tele.Settings{
		Token:  "key",
		Poller: &tele.LongPoller{Timeout: 10 * time.Second},
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		panic(err)
	}

	b.RemoveWebhook(true)

	b.Use(middleware.Recover())
	b.Use(middleware.AutoRespond())

	Handler(b)
	b.Start()
}
