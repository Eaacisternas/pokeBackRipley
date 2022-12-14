package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/*MongoCN es el objeto de conexión a la base de datos*/
var MongoCN = ConectarBD()
var uri = "mongodb+srv://eduardoAguirre:Abril201118@cluster0.ep7djah.mongodb.net/?retryWrites=true&w=majority"

/*ConerctarBD es la función que me permite conectar la BD*/
func ConectarBD() *mongo.Client {
	var clientOptions = options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexión exitosa con la base de datos")
	return client
}

/*valConnectiones hace el ping a la Bd*/

func ValConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}

/*GetCollection obtiene la coleccion de datos desde la base de datos*/
func GetCollection(collection string) *mongo.Collection {
	database := "pokeBack"
	client, err := mongo.NewClient(options.Client().ApplyURI(uri))
	if err != nil {
		panic(err.Error())
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		panic(err.Error())
	}
	return client.Database(database).Collection(collection)
}
