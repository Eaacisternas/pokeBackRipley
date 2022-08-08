package models

//pokemon datos del pokemon
type Pokemon struct {
	Orden  int    `json:"order"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
	Height int    `json:"height"`
}

//pokemons Lista de los pokemones
type Pokemons []Pokemon
