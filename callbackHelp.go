package main

import "fmt"

func CallbackHelp() error {
	fmt.Println("Welcome to trimRL!!")
	fmt.Println("this is the list of commands:")

	availableCommands := GetCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()

	return nil
}
