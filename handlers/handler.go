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

func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/procesarPokemon", middlew.ValidacionDB(routers.RegistrarPokemon)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
