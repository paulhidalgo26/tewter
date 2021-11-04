package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/paulhidalgo26/tewter/bd"
)

//ListaUsiarios leo la lista de los usuarios
func ListaUsiarios(w http.ResponseWriter, r *http.Request) {

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)

	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}
	pag := int64(pagTemp)

	resultado, status := bd.LeerUsuariosTodos(IDUsauario, pag, search, typeUser)

	if !status {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resultado)

}
