package bd

import (
	"context"
	"time"

	"github.com/paulhidalgo26/tewter/models"
)

//InsertoRelacion graba la relacionen la base de datos
func InsertoRelacion(t models.Relacion) (bool, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("relacion")

	_, err := col.InsertOne(ctx, t)

	if err != nil {
		return false, err

	}
	return true, nil

}
