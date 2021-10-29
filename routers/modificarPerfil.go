package routers

import (
	"encoding/json"
	"net/http"

	"github.com/paulhidalgo26/tewter/bd"
	"github.com/paulhidalgo26/tewter/models"
)

//ModificarPerfil modifica el perfil del usario
func ModificarPerfil(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorectos"+err.Error(), 400)
		return
	}
	var status bool
	status, err = bd.ModificoResgistro(t, IDUsauario)
	if err != nil {
		http.Error(w, "Ocurrio un  error al intentar modificar el registro. Reitente nuevamente"+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se a logrado modoficar el registro de usuario "+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
