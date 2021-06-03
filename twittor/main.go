package main

import (
	"log"

	"github.com/webdesigner98everyone/twittor/bd"
	"github.com/webdesigner98everyone/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}
