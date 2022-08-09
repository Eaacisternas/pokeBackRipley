package models

/*Region almacena los datos de la region */
type Region struct {
	Name    string          `json:"name"`
	Pokemon []NumeroPokemon `json:"pokemon_entries,omitempty"`
}

/*NumeroPokemon almacena los datos del numero de poemones encontrados en la region */
type NumeroPokemon struct {
	EntryNo int `json:"entry_number"`
}
