package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/paulhidalgo26/tewter/models"
	"go.mongodb.org/mongo-driver/bson"
)

//ConsultarRelacion consulata si existe relacion en la base de datos
func ConsultarRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}
	var resultado models.Relacion
	fmt.Println("Variable resutado", resultado)
	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		fmt.Println( err.Error())
		return false, err

	}

	return true, nil
}
