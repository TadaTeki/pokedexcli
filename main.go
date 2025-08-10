package main

import (
	"fmt"
	"strings"
)

func main() {
	test_text := "Hello, World!"
	fmt.Println(test_text)
	cleanInput(test_text)
}

func cleanInput(text string) []string {
	words := []string{}
	if text == "" {
		return []string{}
	}
	var sep = " "
	words = strings.Split(text, sep)
	fmt.Println(words)
	return words
}
