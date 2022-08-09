package middlew

import (
	"net/http"

	database "github.com/Eaacisternas/pokeBackRipley/database"
)

func ValidacionDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if database.ValConnection() == 0 {
			http.Error(w, "Conexión perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
