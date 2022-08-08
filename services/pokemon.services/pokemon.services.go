package pokemonservices

import (
	pokemonrepository "pokeBackRipley/repositories/pokemon.repository"

	"../../models"
)

func Create(pokemon models.Pokemon) error {
	err := pokemonrepository.Create(pokemon)
	if err != nil {
		return err
	}
	return nil
}
func Read() (models.Pokemons, error) {
	pokemons, err := pokemonrepository.Read()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}
