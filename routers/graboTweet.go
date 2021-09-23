package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/miguelmalagaortega/twittor/bd"
	"github.com/miguelmalagaortega/twittor/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {

	var mensaje models.Tweet

	err := json.NewDecoder(r.Body).Decode(&mensaje)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar insertar el registro, reintente nuevamente "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Nose ha logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
