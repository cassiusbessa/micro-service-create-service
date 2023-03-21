package errors

import (
	"encoding/json"
	"net/http"
)

func SendError(w http.ResponseWriter, status int, message interface{}) {
	a, _ := json.Marshal(message)
	w.WriteHeader(status)
	w.Write([]byte(a))
}
