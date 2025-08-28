package main

import "fmt"

func CallbackHelp() error {
	fmt.Println("These are the list of commands:")
	fmt.Println()

	availableCommands := GetCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	fmt.Println()

	return nil
}
