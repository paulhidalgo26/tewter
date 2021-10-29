package routers

import (
	"encoding/json"
	"net/http"

	"github.com/paulhidalgo26/tewter/bd"
)

//Verperfil permite extraer los valoresdel perfil
func Verperfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parmetro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Ocurio un error al intentar buscar el registro"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/jsom")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)

}
