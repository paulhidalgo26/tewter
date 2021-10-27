package routers

import (
	"encoding/json"
	"net/http"

	"github.com/paulhidalgo26/tewter/bd"
	"github.com/paulhidalgo26/tewter/models"
)

//Registro es la funcion para crear en  la BD el regstro de usuario
func Registro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}
	if len(t.Password) < 6 {
		http.Error(w, "Debe especificar un contrseña de minimo 6 caracteres", 400)
		return
	}

	_, encontrado, _ := bd.ChequeYaExisteUsuario(t.Email)
	if encontrado == true {
		http.Error(w, "Ya existe un usuario registrado con ese email", 400)
		return
	}

	_, status, err := bd.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar registrar el  usuario"+err.Error(), 400)
		return
	}
	if status == false {
		http.Error(w, "No se a logrado nsertar el registro del usuario"+err.Error(), 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
