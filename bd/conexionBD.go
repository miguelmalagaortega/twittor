package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoCN es la Variable que sera exportada a los demas archivos y es el objeto de conexion
var MongoCN = ConectarBD()

// Variable interna
var clientOptions = options.Client().ApplyURI("mongodb+srv://admin:Dcn123@cluster0.w6vsr.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConectarBD es la funcion que me permite conectar la BD*/
func ConectarBD() *mongo.Client {
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

	log.Println("Conexion exitosa con la BD")
	return client
}

/* ChequeoConnection es el Ping a la BD*/
func ChequeoConnection() int {
	err := MongoCN.Ping(context.TODO(), nil)

	if err != nil {
		return 0
	}

	return 1
}
