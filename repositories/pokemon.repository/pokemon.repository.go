package pokemonrepository

import (
	"context"

	database "github.com/Eaacisternas/pokeBackRipley/database"
	models "github.com/Eaacisternas/pokeBackRipley/models"
)

var collection = database.GetCollection("PokemonCollection")
var ctx = context.Background()

func Create(pokemon models.Pokemon) error {
	var err error
	_, err = collection.InsertOne(ctx, pokemon)
	if err != nil {
		return err
	}
	return nil
}
func Read() (models.Pokemons, error) {
	var pokemons models.Pokemons
	return pokemons, nil
}
