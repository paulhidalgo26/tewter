package routers

import (
	"encoding/json"
	"net/http"

	"github.com/paulhidalgo26/tewter/bd"
	"github.com/paulhidalgo26/tewter/models"
)

//ConsultarRelacion chequea si hay relacion
func ConsultarRelacion(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var t models.Relacion

	t.UsuarioID = IDUsauario
	t.UsuarioRelacionID = ID

	var resp models.RConsultaRelacion

	status, err := bd.ConsultarRelacion(t)

	if err != nil || !status {
		resp.Status = false

	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

}
