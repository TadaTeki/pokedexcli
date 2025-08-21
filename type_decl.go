package main

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

// Define structs to match the JSON structure of the PokeAPI response
type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string  `json:"name"`
		URL  *string `json:"url"`
	} `json:"results"`
}

type config struct {
	Next     string
	Previous string
}
