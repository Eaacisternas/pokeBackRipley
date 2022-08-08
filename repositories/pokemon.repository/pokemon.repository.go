package pokemonrepository

import (
	models "../../models"
)

func Create(pokemon models.Pokemon) error {
	return nil
}
func Read() (models.Pokemons, error) {
	var pokemons models.Pokemons
	return pokemons, nil
}
