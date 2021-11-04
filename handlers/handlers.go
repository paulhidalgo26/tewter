package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/paulhidalgo26/tewter/middlew"
	"github.com/paulhidalgo26/tewter/routers"
	"github.com/rs/cors"
)

//Manejadores seteo mi puerto , el handelr y pongo a escuhar mi servidor
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.Verperfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.BorrarTweet))).Methods("DELETE")

	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	router.HandleFunc("/crearRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.CrearRelacion))).Methods("POST")
	router.HandleFunc("/eliminarRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultarRelacion))).Methods("GET")

	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsiarios))).Methods("GET")
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeerTweetRelacion))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
