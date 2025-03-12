package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	nextUrl string
	prevUrl string
}

func initializeRepl() {
	cfg := &Config{
		nextUrl: "https://pokeapi.co/api/v2/location-area/",
		prevUrl: "",
	}

	cliCommands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display next 20 area locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display previous 20 area locations",
			callback:    commandMapb,
		},
	}
	cliCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp(cliCommands),
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cleaned := cleanInput(input)

		if len(cleaned) > 0 {
			commandName := cleaned[0]
			if command, exists := cliCommands[commandName]; exists {
				err := command.callback(cfg)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}
		} else {
			fmt.Println("No command entered")
		}
	}
}

func cleanInput(text string) []string {
	fixed_str := strings.Fields(strings.ToLower(text))

	return fixed_str
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) func(*Config) error {
	return func(*Config) error {
		fmt.Println("Welcome to the Pokedex!")
		fmt.Println("Usage:")
		fmt.Println()
		for _, cmd := range commands {
			fmt.Printf("%s: %s\n", cmd.name, cmd.description)
		}
		return nil
	}
}

type LocationAreaResponse struct {
	Count   int            `json:"count"`
	Next    *string        `json:"next"`
	Prev    *string        `json:"previous"`
	Results []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func commandMap(cfg *Config) error {
	// check if last page before trying fetch
	if cfg.nextUrl == "" {
		fmt.Println("You have reached the last page")
	} else {
		locationResp, err := fetchLocationData(cfg.nextUrl)
		if err != nil {
			return fmt.Errorf("Error fetching next location data: %v", err)
		}

		//iterate over location response, print
		for _, loc := range locationResp.Results {
			fmt.Println(loc.Name)
		}

		if locationResp.Next != nil {
			cfg.nextUrl = *locationResp.Next
		} else {
			cfg.nextUrl = ""
		}
		if locationResp.Prev != nil {
			cfg.prevUrl = *locationResp.Prev
		} else {
			cfg.prevUrl = ""
		}
	}
	return nil
}

func commandMapb(cfg *Config) error {
	// check if first page before trying fetch
	if cfg.prevUrl == "" {
		fmt.Println("you're on the first page")
	} else {
		locationResp, err := fetchLocationData(cfg.prevUrl)
		if err != nil {
			return fmt.Errorf("Error fetching previous location data: %v", err)
		}

		for _, loc := range locationResp.Results {
			fmt.Println(loc.Name)
		}

		if locationResp.Next != nil {
			cfg.nextUrl = *locationResp.Next
		} else {
			cfg.nextUrl = ""
		}
		if locationResp.Prev != nil {
			cfg.prevUrl = *locationResp.Prev
		} else {
			cfg.prevUrl = ""
		}
	}
	return nil
}

// helper function for both map funcs
func fetchLocationData(url string) (*LocationAreaResponse, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch response: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if res.StatusCode > 299 {
		return nil, fmt.Errorf("Response failed with status code: %d", res.StatusCode)
	}
	if err != nil {
		return nil, fmt.Errorf("Failed to read body: %v", err)
	}

	var locationResp LocationAreaResponse
	err = json.Unmarshal(body, &locationResp)
	if err != nil {
		return nil, fmt.Errorf("Failed to unmarshal: %v", err)
	}

	return &locationResp, nil
}
