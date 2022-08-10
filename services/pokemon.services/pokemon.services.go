package pokemonservices

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	models "github.com/Eaacisternas/pokeBackRipley/models"
	pokemonrepository "github.com/Eaacisternas/pokeBackRipley/repositories/pokemon.repository"
)

/*Create, invoca al repositorio que almacena un registro en la base de datos*/
func Create(pokemon models.Pokemon) error {
	err := pokemonrepository.Create(pokemon)
	if err != nil {
		return err
	}
	return nil
}

/*Read, invoca al repositorio que lee los registros en la base de datos*/
func Read() error {
	err := pokemonrepository.Read()
	if err != nil {
		return err
	}
	return nil
}

/*ListarPokemon, Lee la pokeapi(https://pokeapi.co/) y procesa los pokemons deacuerdo a lo requerido*/
func ListarPokemon(w http.ResponseWriter, r *http.Request) error {
	var pokemonObject models.Pokemon
	var pokemons models.Pokemons
	for i := 1; i <= 10; i++ {
		response, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + fmt.Sprint(i))
		if err != nil {
			return err
		}
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		json.Unmarshal(responseData, &pokemonObject)

		pokemons = append(pokemons, &pokemonObject)
		Create(pokemonObject)
	}
	w.Header().Set("contect-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(pokemons)

	return nil
}
