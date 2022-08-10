package routers

import (
	"net/http"

	pokemonrepository "github.com/Eaacisternas/pokeBackRipley/repositories/pokemon.repository"
	pokemonservices "github.com/Eaacisternas/pokeBackRipley/services/pokemon.services"
)

// RegistrarPokemon datos enviados por verbo POST
func RegistrarPokemon(w http.ResponseWriter, r *http.Request) {
	err := pokemonrepository.Read(w, r)
	if err != nil {
		http.Error(w, "Error en el registro de los datos recibidos "+err.Error(), 400)
		return
	}
}

// ListarPokemon datos enviados por verbo GET
func ListarPokemon(w http.ResponseWriter, r *http.Request) {
	err := pokemonservices.ListarPokemon(w, r)
	if err != nil {
		http.Error(w, "Error en el listado de los pokemon", 400)
		return
	}
}
