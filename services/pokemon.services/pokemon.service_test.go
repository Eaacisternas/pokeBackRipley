package pokemonservices_test

import (
	"testing"

	models "github.com/Eaacisternas/pokeBackRipley/models"
	pokemonservices "github.com/Eaacisternas/pokeBackRipley/services/pokemon.services"
)

/*TestCreate, test unitario para Create de pokemonservice*/
func TestCreate(t *testing.T) {
	pokemon := models.Pokemon{
		Orden:  1,
		Name:   "Bulbasur",
		Weight: 67,
		Height: 7,
	}
	err := pokemonservices.Create(pokemon)
	if err != nil {
		t.Error("La prueba de persistencia de los datos de pokemon ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}

/*TestRead, test unitario para Read de pokemonservice*/
func TestRead(t *testing.T) {
	pokemons, err := pokemonservices.Read()
	if err != nil {
		t.Error("La prueba de persistencia de los datos de pokemon ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
	if len(pokemons) == 0 {
		t.Error("La consulta no ha traido datos")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}

/*TestListarPokemon, test unitario para ListarPokemon de pokemonservice*/
func TestListarPokemon(t *testing.T) {
	err := pokemonservices.ListarPokemon("kanto")
	if err != nil {
		t.Error("La prueba de persistencia de los datos de pokemon por region ha fallado ")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}

/*TestCrearArchivo, test unitario para CrearArchivo de pokemonservice*/
func TestCrearArchivo(t *testing.T) {
	pokemon := models.Pokemon{
		Orden:  2,
		Name:   "Ivisur",
		Weight: 67,
		Height: 7,
	}
	err := pokemonservices.CrearArchivo(pokemon, "kanto")
	if err != nil {
		t.Error("La prueba de crear un archivo ha fallado")
		t.Fail()
	} else {
		t.Log("La prueba finalizo con exito")
	}
}
