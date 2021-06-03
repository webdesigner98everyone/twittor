package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/webdesigner98everyone/twittor/middlew"
	"github.com/webdesigner98everyone/twittor/routers"
)

/* Manejadores seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	//Rutas

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	//pone a escuchar el puerto
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
