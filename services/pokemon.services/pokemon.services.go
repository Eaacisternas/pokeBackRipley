package pokemonservices

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
func ListarPokemon() error {
	var pokemonObject models.Pokemon
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err.Error())
	}
	var responseObject models.Region
	json.Unmarshal(responseData, &responseObject)
	for i := 1; i <= len(responseObject.Pokemon); i++ {
		response, err = http.Get("http://pokeapi.co/api/v2/pokemon/" + fmt.Sprint(i))
		if err != nil {
			fmt.Print(err.Error())
			os.Exit(1)
		}
		responseData, err = io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err.Error())
		}
		json.Unmarshal(responseData, &pokemonObject)
		fmt.Println(pokemonObject)
		Create(pokemonObject)
	}
	return nil
}
