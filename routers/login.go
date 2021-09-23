package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/miguelmalagaortega/twittor/bd"
	"github.com/miguelmalagaortega/twittor/jwt"
	"github.com/miguelmalagaortega/twittor/models"
)

func Login(w http.ResponseWriter, r *http.Request) {
	// Enviamos a la cabecera el formato tipo json
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	// resivimos los dos datos, email y password por medio del Body luego lo decodificamos a json y los guardamos en t
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		// devolvemos el error en caso halla
		http.Error(w, "Usuario y/o contraseña invalida "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	// llamamos a la funcion IntentoLogin para ver si nos devulve un usuario o un error
	documento, existe := bd.IntentoLogin(t.Email, t.Password)

	if existe == false {
		http.Error(w, "Usuario y/o contraseña invalida", 400)
		return
	}

	// esto devolvera el token o el error
	jwtKey, err := jwt.GeneroJWT(documento)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el token correspondiente "+err.Error(), 400)
		return
	}

	// Armamos un json con el token para luego devolverlo al navegador
	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	// Con esto devolveremos el token al navegador
	w.Header().Set("content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	// Grabar una cookie
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
