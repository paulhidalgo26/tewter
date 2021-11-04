package routers

import (
	"net/http"

	"github.com/paulhidalgo26/tewter/bd"
)

func BorrarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parametro ID", http.StatusBadRequest)
		return

	}
	err := bd.EliminarTweet(ID, IDUsauario)

	if err != nil {
		http.Error(w, "Ocurio un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
