package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/paulhidalgo26/tewter/models"
)

//GeneroJWT genera el encriptado con JWT
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("Secret")
	payload := jwt.MapClaims{
		"email":           t.Email,
		"nombre":          t.Nombre,
		"apellidos":       t.Apellidos,
		"fechaNacimiento": t.FechaNacimiento,
		"biografia":       t.Biografia,
		"ubicacion":       t.Ubicacion,
		"sitioweb":        t.SitioWeb,
		"_id":             t.ID.Hex(),
		"exp":             time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	tokenstr, err := token.SignedString(miClave)
	if err != nil {
		return "error de token  " + err.Error(), err
	}
	return tokenstr, nil
}
