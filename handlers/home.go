package handlers

// Encargado de procesar la petición de la ruta principal

import (
	"encoding/json"
	"net/http"
	"rest-api-go/server"
)

type HomeResponse struct {
	Message string `json:"mensage"` //En go la ponemos en mayusculas, cuando se serialice a json se ponde message en minúsculas
	Status  bool   `json:"status"`
	//StatusCode int16  `json:"statuscode"`
}

func HomeHandler(s server.Server) http.HandlerFunc {
	//w -> encarga de dar la respuesta al cliente
	//r -> Data que envía el cliente, la tomamos y hacemos los procesamientos.
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json") //Le digo que la respuesta es JSON
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(HomeResponse{
			Message: "Tu primera REST API",
			Status:  true,
		})

	}
}
