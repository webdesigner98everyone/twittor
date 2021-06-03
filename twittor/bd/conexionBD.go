package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN  es el objeto de conexcion a la bd*/

var MongoCN = ConectarBD()

//colocamos un parametro que es la URL de la BD
var clientOptions = options.Client().ApplyURI("mongodb+srv://root:root@cluster0.1xrlq.mongodb.net/admin")

/* ConectarBD es la funcion que me permite conectar la BD*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	//es para validar si esta o no disponible la bd
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa con la BD")
	return client
}

/* ChequeoConnection es el ping a la bd*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
