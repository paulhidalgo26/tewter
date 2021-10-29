package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/paulhidalgo26/tewter/bd"
	"github.com/paulhidalgo26/tewter/models"
)

//GraboTweet permite grabar el tweet en la base de datos
func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "tweet invalido"+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsauario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, errr := bd.InsertoTweet(registro)

	if errr != nil {
		http.Error(w, "Ocurrio un error al intentar insertar un registro, reitente nuevamente "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "no se a logrado insertar en tweet ", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
