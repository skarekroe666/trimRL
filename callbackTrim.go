package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
)

func CallbackTrim() error {
	fmt.Print("Enter your url: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	lowered := strings.ToLower(text)
	cleanLink := cleanString(lowered)
	hash := generateHash(cleanLink)

	link := "https://trim.rl/" + hash

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, text, http.StatusFound)
	})

	fmt.Printf("here's your trimmed url: %s\n", link)
	fmt.Println()

	return nil
}

func generateHash(str string) string {
	l := 5
	hash := make([]byte, l)
	for i := range hash {
		hash[i] = str[rand.Intn(len(str))]
	}

	return string(hash)
}

func cleanString(str string) string {
	result := ""
	for v := range str {
		if str[v] == '/' || str[v] == ':' || str[v] == '?' || str[v] == '.' || str[v] == '=' || str[v] == '-' {
			v++
		}
		result = result + string(str[v])
	}

	return result
}
