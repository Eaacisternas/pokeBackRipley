package models

type Region struct {
	Name    string          `json:"name"`
	Pokemon []NumeroPokemon `json:"pokemon_entries"`
}

type NumeroPokemon struct {
	EntryNo int `json:"entry_number"`
}
