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
func Read() (models.Pokemons, error) {
	pokemons, err := pokemonrepository.Read()
	if err != nil {
		return nil, err
	}
	return pokemons, nil
}

/*ListarPokemon, Lee la pokeapi(https://pokeapi.co/) y procesa los pokemons deacuerdo a lo requerido*/
func ListarPokemon(region string) error {
	var pokemonObject models.Pokemon
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/" + region)
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
		Create(pokemonObject)
	}
	return nil
}

/*CrearArchivo, Crea un archivo csv en documents, con la data extraida desde la base de datos*/
func CrearArchivo(pokemon models.Pokemon, region string) error {
	filename := "../../documents/" + region + ".csv"
	linea := fmt.Sprint(pokemon.Orden) + ";" + string(pokemon.Name) + ";" + fmt.Sprint(pokemon.Weight) + ";" + fmt.Sprint(pokemon.Height)
	if !archivoExiste(filename) {
		archivo, err := os.Create(filename)
		if err != nil {
			fmt.Print("Hubo un error al crear el archivo")
			return nil
		}
		fmt.Fprintln(archivo, linea)
		archivo.Close()
	} else {
		archivo, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Hubo un error al modificar el archivo")
			return nil
		}
		_, err = archivo.WriteString("\n" + linea)
		if err != nil {
			fmt.Println("Hubo un error al ingresar el texto en el archivo")
		}
	}

	return nil
}

/*archivoExiste, valida que exista el archivo en la base de datos*/
func archivoExiste(ruta string) bool {
	if _, err := os.Stat(ruta); os.IsNotExist(err) {
		return false
	}
	return true
}
