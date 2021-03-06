package routers

import (
	"net/http"

	"github.com/paulhidalgo26/tewter/bd"
	"github.com/paulhidalgo26/tewter/models"
)

//AltaRelacion realiza el regrsitro de la relacion etre usuarios
func CrearRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parametro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsauario
	t.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se a logrado   insertar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
