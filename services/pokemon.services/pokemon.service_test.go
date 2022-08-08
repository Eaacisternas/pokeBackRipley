package pokemonservices_test

import (
	"testing"

	models "github.com/Eaacisternas/pokeBackRipley/models"
	pokemonservices "github.com/Eaacisternas/pokeBackRipley/services/pokemon.services"
)

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
func TestRead(t *testing.T) {

}
