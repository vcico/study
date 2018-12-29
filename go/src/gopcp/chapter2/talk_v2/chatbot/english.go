package chatbot

import (
	"fmt"
	"strings"
)

type simpleEN struct {
	name string
	talk Talk
}

func NewSimpleEn(name string, talk Talk) Chatbot {
	return &simpleEN{
		name: name,
		talk: talk,
	}
}
