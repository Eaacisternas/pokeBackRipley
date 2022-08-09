package main

import (
	repository "github.com/Eaacisternas/pokeBackRipley/repositories/pokemon.repository"
	// "github.com/Eaacisternas/pokeBackRipley/handlers"
)

/*Funcion principal del programa*/
func main() {
	// handlers.Manejadores()
	repository.Read()

}
