package main

import (
	"log"

	"github.com/paulhidalgo26/tewter/bd"

	"github.com/paulhidalgo26/tewter/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()

}
