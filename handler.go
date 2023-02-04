package main

import (
	"fmt"

	tele "github.com/3JoB/telebot"
)

var (
	UserStatus map[int64]int
	UserName   map[int64]string
	UserBirth  map[int64]string
)

const (
	StatusNone = iota
	StatusStep1
	StatusStep2
)

func init() {
	UserStatus = make(map[int64]int)
	UserName = make(map[int64]string)
	UserBirth = make(map[int64]string)
}

func Handler(b *tele.Bot) {
	b.Handle("/start", Start)
	b.Handle(tele.OnText, Text)
}

func Text(c tele.Context) error {
	switch UserStatus[c.Chat().ID] {
	case StatusStep1:
		return start1(c)
	case StatusStep2:
		return start2(c)
	}
	return nil
}

func Start(c tele.Context) error {
	c.Send("Please enter a nickname!")
	UserStatus[c.Chat().ID] = StatusStep1
	return nil
}

func start1(c tele.Context) error {
	c.Send("Please enter birthday!")
	UserName[c.Chat().ID] = c.Message().Text
	UserStatus[c.Chat().ID] = StatusStep2
	return nil
}

func start2(c tele.Context) error {
	UserBirth[c.Chat().ID] = c.Message().Text
	c.Send(fmt.Sprintf("Here is your message:\nName: %v \nBirthday: %v", UserName[c.Chat().ID], UserBirth[c.Chat().ID]))
	// This must be released after completion, or a /cancel command can be created to allow the user to manually release
	delete(UserName, c.Chat().ID)
	delete(UserBirth, c.Chat().ID)
	return nil
}
