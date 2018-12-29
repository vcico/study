package main

import (
	"flag"
	"fmt"
)

// 对话使用的聊天机器人
var chatbotName string

func init() {
	flag.StringVar(&chatbotName, "chatbot", "simple.en", "The chatbot's name for dialogue")
}

func main() {
	flag.Parse()
	chatbot.Register(chatbot.NewSimpleEn("simple.en", nil))
}
