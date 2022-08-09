package routers

import (
	"encoding/json"
	"net/http"

	models "github.com/Eaacisternas/pokeBackRipley/models"
	pokemonrepository "github.com/Eaacisternas/pokeBackRipley/repositories/pokemon.repository"
)

// RegistrarPokemon datos enviados por verbo http
func RegistrarPokemon(w http.ResponseWriter, r *http.Request) {
	var region models.Region
	err := json.NewDecoder(r.Body).Decode(&region)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	err = pokemonrepository.Read()
	if err != nil {
		http.Error(w, "Error en el registro de los datos recibidos", 400)
		return
	}
}
