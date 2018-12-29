package main

import (
	"bufio"
	"fmt"
	"os"

	//	"reflect"
	"strings"
)

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please input your name:")
	input, err := inputReader.ReadString('\n')
	if err != nil {
		fmt.Printf("An error occurred: %s \n", err)
		os.Exit(1) // 异常退出
	}
	name := input[:len(input)-1]
	fmt.Printf("--%s--%s--", input, name)
	fmt.Printf("Hello %s what can i do for you ?\n", name)

	for {
		input, err = inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred : %s\n", err)
			continue
		}
		input = input[:len(input)-2]
		input = strings.ToLower(input)
		switch input {
		case "":
			continue
		case "nothing", "bye":
			fmt.Println("bye!")
			os.Exit(0)
		default:
			fmt.Println("Sorry , i didnt catch you ")
		}
	}

}
