package middlew

import (
	"net/http"

	"github.com/paulhidalgo26/tewter/bd"
)

//ChequeoBD es el middelw que me permite conecer el estado de la BD
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la base de datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
