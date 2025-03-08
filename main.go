package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	fixed_str := strings.Fields(strings.ToLower(text))

	return fixed_str
}
