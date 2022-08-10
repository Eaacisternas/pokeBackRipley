package handlers

import (
	"log"
	"net/http"
	"os"

	middlew "github.com/Eaacisternas/pokeBackRipley/middlewares"
	routers "github.com/Eaacisternas/pokeBackRipley/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores evalua la accion sobre la peticion http, y envia deacuerdo al Verbo y los parametros que este tenga.*/
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/procesarPokemon", middlew.ValidacionDB(routers.RegistrarPokemon)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "7070"
	}
	router.HandleFunc("/listarPokemon", routers.ListarPokemon).Methods("GET")
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
