package middlew

import (
	"net/http"

	"github.com/miguelmalagaortega/twittor/bd"
)

// los middlewares reciben algo y devuelven lo mismo
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConnection() == 0 {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
