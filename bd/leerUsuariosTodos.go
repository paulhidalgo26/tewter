package bd

import (
	"context"
	"fmt"
	"time"

	"github.com/paulhidalgo26/tewter/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeerUsuariosTodos lee los usuarios registrados en el sistema, si se recibe "R" lee los relacionados con el usuario
func LeerUsuariosTodos(ID string, page int64, search string, tipo string) ([]*models.Usuario, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("twiter")
	col := db.Collection("usuarios")

	var resultado []*models.Usuario

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	query := bson.M{
		"nombre": bson.M{"$regex": `(?i)` + search},
	}

	cursor, err := col.Find(ctx, query, findOptions)
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}

	var encontrado, incluir bool

	for cursor.Next(ctx) {
		var s models.Usuario
		err1 := cursor.Decode(&s)
		if err1 != nil {
			fmt.Println(err.Error())
			return resultado, false
		}
		var r models.Relacion
		r.UsuarioID = ID
		r.UsuarioRelacionID = s.ID.Hex()

		incluir = false

		encontrado, err = ConsultarRelacion(r)

		if tipo == "new" && !encontrado {
			incluir = true
		}
		if tipo == "follow" && encontrado {
			incluir = true
		}

		if r.UsuarioRelacionID == ID {
			incluir = false
		}

		if incluir {
			s.Password = ""
			s.Biografia = ""
			s.SitioWeb = ""
			s.Ubicacion = ""
			s.Banner = ""
			s.Email = ""
			resultado = append(resultado, &s)
		}
	}

	err = cursor.Err()
	if err != nil {
		fmt.Println(err.Error())
		return resultado, false
	}
	cursor.Close(ctx)
	return resultado, true

}
