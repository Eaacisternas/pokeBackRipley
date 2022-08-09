package routers

import (
	"encoding/json"
	"net/http"

	models "github.com/Eaacisternas/pokeBackRipley/models"
	pokemonservices "github.com/Eaacisternas/pokeBackRipley/services/pokemon.services"
)

func RegistrarPokemon(w http.ResponseWriter, r *http.Request) {
	var region models.Region
	err := json.NewDecoder(r.Body).Decode(&region)
	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}
	err = pokemonservices.ListarPokemon(string(region.Name))
	if err != nil {
		http.Error(w, "Error en el registro de los datos recibidos", 400)
		return
	}
}
