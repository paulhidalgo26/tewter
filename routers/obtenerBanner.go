package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/paulhidalgo26/tewter/bd"
)

//ObtenerAvatar envia el avatar al HTTP
func ObtenerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID ", http.StatusBadRequest)
		return
	}
	perfil, err := bd.BuscoPerfil(ID)

	if err != nil {

		http.Error(w, "Usuario No Encontrado ", http.StatusBadRequest)
		return
	}

	OpenFile, err1 := os.Open("uploads/banners" + perfil.Banner)

	if err1 != nil {
		http.Error(w, "Imagen no encontrada ", http.StatusBadRequest)
		return
	}
	_, err2 := io.Copy(w, OpenFile)
	if err2 != nil {
		http.Error(w, "Error al copiar imagen ", http.StatusBadRequest)
		return
	}

}
