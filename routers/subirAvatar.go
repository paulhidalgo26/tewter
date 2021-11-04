package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/paulhidalgo26/tewter/bd"
	"github.com/paulhidalgo26/tewter/models"
)

//SubirAvatar sube el avatar al servidor
func SubirAvatar(w http.ResponseWriter, r *http.Request) {

	file, handler, err := r.FormFile("avatar")
	var extension = strings.Split(handler.Filename, ".")[1]

	var archivo string = "uploads/avatars/" + IDUsauario + "." + extension

	f, err1 := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)

	if err1 != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err2 := io.Copy(f, file)
	if err2 != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool
	usuario.Avatar = IDUsauario + "." + extension
	status, err3 := bd.ModificoResgistro(usuario, IDUsauario)
	if err3 != nil || !status {
		http.Error(w, "Error al grabar avatar ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}
