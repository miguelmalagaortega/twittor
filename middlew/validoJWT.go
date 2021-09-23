package middlew

import (
	"net/http"

	"github.com/miguelmalagaortega/twittor/routers"
)

func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// leemos de la cabecera el valor Authorization y lo mandamos a ProcesoToken
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authorization"))

		if err != nil {
			http.Error(w, "Error en el token! "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)

	}

}
