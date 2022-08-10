package pokemonrepository

import (
	"context"
	"fmt"
	"net/http"
	"os"

	database "github.com/Eaacisternas/pokeBackRipley/database"
	models "github.com/Eaacisternas/pokeBackRipley/models"
	ftpservices "github.com/Eaacisternas/pokeBackRipley/services/ftp.services"

	"go.mongodb.org/mongo-driver/bson"
)

var collection = database.GetCollection("PokemonCollection")
var ctx = context.Background()

/*Create, almacena los registros en la base de datos*/
func Create(pokemon models.Pokemon) error {
	var err error
	_, err = collection.InsertOne(ctx, pokemon)
	if err != nil {
		return err
	}
	return nil
}

/*Read, lee los registros en la base de datos*/
func Read(w http.ResponseWriter, r *http.Request) error {
	filter := bson.D{}
	cur, err := collection.Find(ctx, filter)
	if err != nil {
		return err
	}
	for cur.Next(ctx) {
		var pokemon models.Pokemon
		err = cur.Decode(&pokemon)
		if err != nil {
			return err
		}
		CrearArchivo(pokemon)
	}
	err = ftpservices.SubirArchivo()
	if err != nil {
		return err
	}

	w.Header().Set("contect-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	return nil
}

/*CrearArchivo, Crea un archivo csv en documents, con la data extraida desde la base de datos*/
func CrearArchivo(pokemon models.Pokemon) error {
	filename := "./documents/kanto.csv"
	linea := fmt.Sprint(pokemon.Orden) + ";" + string(pokemon.Name) + ";" + fmt.Sprint(pokemon.Weight) + ";" + fmt.Sprint(pokemon.Height) + ";" + pokemon.Sprites.Front_default
	for j := 0; j < len(pokemon.Types); j++ {
		linea = linea + ";" + string(pokemon.Types[j].Tipo.Name)
	}
	if !archivoExiste(filename) {
		archivo, err := os.Create(filename)
		if err != nil {
			fmt.Print("Hubo un error al crear el archivo" + err.Error())
			return err
		}
		fmt.Fprintln(archivo, linea)
	} else {
		archivo, err := os.OpenFile(filename, os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			fmt.Println("Hubo un error al modificar el archivo")
			return err
		}
		_, err = archivo.WriteString("\n" + linea)
		if err != nil {
			fmt.Println("Hubo un error al ingresar el texto en el archivo")
			return err
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
