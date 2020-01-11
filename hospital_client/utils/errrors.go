package utils

import (
	"net/http"
)

func InternalServerError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(" Oops! Internal Server Error Has Occurred:("))
}
func StatusNotFoundError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("Oops! Status Not Found. :("))
}
