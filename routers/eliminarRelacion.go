package routers

import (
	"net/http"

	"github.com/paulhidalgo26/tewter/bd"
	"github.com/paulhidalgo26/tewter/models"
)

func EliminarRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relacion

	t.UsuarioID = IDUsauario
	t.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(t)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se a logrado   borrar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
