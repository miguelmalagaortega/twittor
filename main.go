package main

import (
	"log"
	// para poder llamar archivos de otras rutas debemos usar
	// github.com/miguelmalagaortega/twittor/ con esto ya estamos en la ruta del proyecto
	"github.com/miguelmalagaortega/twittor/bd"
	"github.com/miguelmalagaortega/twittor/handlers"
)

func main() {

	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()

}
