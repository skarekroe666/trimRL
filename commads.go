package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type InputCmds struct {
	Name        string
	Description string
	Callback    func() error
}

func Commands() {
	fmt.Println("Welcome to trimRL!!")
	fmt.Println()
	// fmt.Println("Enter your url to shorten it:")

	for {

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print(">: ")
		scanner.Scan()
		text := scanner.Text()

		input := cleanInput(text)
		if len(input) == 0 {
			continue
		}
		commandName := input[0]

		avalaibleCmds := GetCommands()
		command, ok := avalaibleCmds[commandName]
		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		err := command.Callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func GetCommands() map[string]InputCmds {
	return map[string]InputCmds{
		"help": {
			Name:        "help",
			Description: "prints all the commands",
			Callback:    CallbackHelp,
		},
		"trim": {
			Name:        "trim",
			Description: "this command will trim your long url",
			Callback:    CallbackTrim,
		},
		"exit": {
			Name:        "exit",
			Description: "enter 'exit'to quit",
			Callback:    CallbackExit,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)

	return words
}
