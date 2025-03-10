package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func initializeRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) > 0 {
			fmt.Println("Your command was:", cleaned[0])
		} else {
			fmt.Println("No command entered.")
		}
	}
}

func cleanInput(text string) []string {
	fixed_str := strings.Fields(strings.ToLower(text))

	return fixed_str
}
