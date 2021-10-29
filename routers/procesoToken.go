package routers

import (
	"errors"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/paulhidalgo26/tewter/bd"
	"github.com/paulhidalgo26/tewter/models"
)

//Email valor del email usado en todos los enpoins
var Email string

//IDUsuario es el id delvuelto del modelo que se usara en todos los Endpoints
var IDUsauario string

//ProcesoToken proceso para exraer valores
func ProcesoToken(tk string) (*models.Clain, bool, string, error) {
	miClave := []byte("Secret")
	claims := &models.Clain{}
	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("fromato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, encontrado, _ := bd.ChequeYaExisteUsuario(claims.Email)
		if encontrado {
			Email = claims.Email
			IDUsauario = claims.ID.Hex()
		}
		return claims, encontrado, IDUsauario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err

}
