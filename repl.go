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

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	var cfg config = config{
		Next:     "",
		Previous: "",
	}

	fmt.Print("Welcome to the Pokedex!\n")

	for {
		fmt.Print("Pokedex >")
		if scanner.Scan() {
			if scanner.Text() == "" {
				continue
			}
			arr_input_txt := cleanInput(scanner.Text())

			cmd, ok := supported_cmd[arr_input_txt[0]]
			if !ok {
				fmt.Println("Unknown command")
				continue
			}

			if err := cmd.callback(&cfg); err != nil {
				fmt.Printf("Error executing command '%s': %v\n", cmd.name, err)
			}

		}

	}
}

func cleanInput(text string) []string {

	if text == "" {
		return []string{}
	}
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, "　", " ")
	//	re := regexp.MustCompile(`[.,!?;:"'()]`)
	//	text = re.ReplaceAllString(text, "")

	words := strings.Fields(text)
	return words
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)

	return nil
}

func commandHelp(cfg *config) error {
	fmt.Print("Welcome to the Pockedex!\nUsage:\n\n")

	for key, value := range supported_cmd {
		fmt.Printf("%s: %s\n", key, value.description)
	}

	return nil
}

func commandMap(cfg *config) error {
	var tgt_url string = "https://pokeapi.co/api/v2/location-area/"

	if cfg.Next == "" && cfg.Previous != "" {
		fmt.Println("No next page available")
		return nil
	}
	// If cfg.Next is not empty, use it as the target URL
	if cfg.Next != "" {
		tgt_url = cfg.Next
	}

	//	fmt.Println(tgt_url)
	areas, err := get_LocationArea(tgt_url)
	if err != nil {
		return err
	}

	update_config_LocationArea(cfg, areas) // Update the config with the next and previous URLs

	for _, r := range areas.Results {
		fmt.Println(r.Name)
	}

	return nil
}

func commandMapb(cfg *config) error {
	var tgt_url string = ""

	if cfg.Previous == "" {
		fmt.Println("you're on the first page")
		return nil
	} else {
		tgt_url = cfg.Previous
	}

	//	fmt.Println(tgt_url)
	areas, err := get_LocationArea(tgt_url)
	if err != nil {
		return err
	}
	update_config_LocationArea(cfg, areas) // Update the config with the next and previous URLs

	for _, r := range areas.Results {
		fmt.Println(r.Name)
	}

	return nil
}

func get_LocationArea(tgt_url string) (LocationArea, error) {
	var areas LocationArea = LocationArea{}
	res, err := http.Get(tgt_url)
	if err != nil {
		return areas, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return areas, err
	}

	if err := json.Unmarshal(data, &areas); err != nil {
		return areas, err
	}
	return areas, nil
}

func update_config_LocationArea(cfg *config, areas LocationArea) error {
	if areas.Next != nil {
		cfg.Next = *areas.Next
	} else {
		cfg.Next = ""
	}
	if areas.Previous != nil {
		cfg.Previous = *areas.Previous
	} else {
		cfg.Previous = ""
	}
	return nil
}
