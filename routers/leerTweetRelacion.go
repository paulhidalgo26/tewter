package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/paulhidalgo26/tewter/bd"
)

//LeerTweetRelacion lee los tweets de todos nuestros seguidores
func LeerTweetRelacion(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "Debe enviar el parametro pagina", http.StatusBadRequest)
		return
	}

	pagina, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "Debe enviar el parametro pagina como entero mayor a 0", http.StatusBadRequest)
		return
	}

	respuesta, correcto := bd.LeerTweetSeguidores(IDUsauario, pagina)

	if !correcto {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(respuesta)

}
