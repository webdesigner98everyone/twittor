package middlew

import (
	"net/http"

	"github.com/webdesigner98everyone/twittor/bd"
)

/*chequeoBD es el Middlew que me permite conocer el estado de la bD*/

func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la BD", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
