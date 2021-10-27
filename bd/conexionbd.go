package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//MongoCN objeto de conexion a la base dee datos
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://paul:Wa3tc5govZiZl1gD@cluster0.yz5dc.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

//ConectarBD permite a la base de datos mongo
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
	log.Println("Conexion Exitosa con la BD")
	return client
}

//ChequeoConnection es el ping a la base de datos
func ChequeoConnection() int {

	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {

		return 0
	}
	return 1
}
