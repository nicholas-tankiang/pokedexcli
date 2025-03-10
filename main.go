package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)
		fmt.Println("Your command was:", cleaned[0])
	}
}

func cleanInput(text string) []string {
	fixed_str := strings.Fields(strings.ToLower(text))

	return fixed_str
}
