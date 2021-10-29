package bd

import (
	"context"
	"time"

	"github.com/paulhidalgo26/tewter/models"
	"go.mongodb.org/mongo-driver/bson"
)

// ChequeYaExisteUsuario recibe un email de parametro y chequea si ya existe en la base de datos
func ChequeYaExisteUsuario(email string) (models.Usuario, bool, string) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("usuarios")

	condicion := bson.M{"email": email}

	var resultado models.Usuario

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	ID := resultado.ID.Hex()
	if err != nil {
		return resultado, false, ID
	}
	return resultado, true, ID
}
