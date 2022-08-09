package models

type Region struct {
	Name    string          `json:"name"`
	Pokemon []NumeroPokemon `json:"pokemon_entries,omitempty"`
}

type NumeroPokemon struct {
	EntryNo int `json:"entry_number"`
}
