package bd

import (
	"context"
	"time"

	"github.com/paulhidalgo26/tewter/models"
)

//BorroRelacionborrar la relacion de la base de datos
func BorroRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("relacion")

	_, err := col.DeleteOne(ctx, t)
	if err != nil {
		return false, err
	}
	return true, nil
}
