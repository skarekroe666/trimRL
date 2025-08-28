package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CallbackTrim() error {
	fmt.Print("Enter your url: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	lowered := strings.ToLower(text)

	fmt.Printf("here's your trimmed url: %s\n", lowered)
	fmt.Println()

	return nil
}
