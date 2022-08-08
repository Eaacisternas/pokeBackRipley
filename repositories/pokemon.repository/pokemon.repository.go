package pokemonrepository

import (
	"context"
	"fmt"

	database "github.com/Eaacisternas/pokeBackRipley/database"
	models "github.com/Eaacisternas/pokeBackRipley/models"
	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("PokemonCollection")
var ctx = context.Background()

/*Create, almacena los registros en la base de datos*/
func Create(pokemon models.Pokemon) error {
	var err error
	_, err = collection.InsertOne(ctx, pokemon)
	if err != nil {
		return err
	}
	return nil
}

/*Read, lee los registros en la base de datos*/
func Read() (models.Pokemons, error) {
	var pokemons models.Pokemons
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	for cur.Next(ctx) {
		var pokemon models.Pokemon
		err = cur.Decode(&pokemon)
		if err != nil {
			return nil, err
		}
		fmt.Println(pokemon)
		pokemons = append(pokemons, &pokemon)
	}
	return pokemons, nil
}
