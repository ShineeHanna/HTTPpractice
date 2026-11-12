package pkg

import (
	"encoding/json"
	"net/http"
)

func Json(w http.ResponseWriter, data any, statusCode int) {
	w.Header().Set("Content-Type", "application/json") //обозначем каким типом у нас будет ответ
	w.WriteHeader(statusCode)                          //обозначем статус код
	json.NewEncoder(w).Encode(data)
}
