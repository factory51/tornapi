package responses

import (
	"net/http"
)

func StandandResponseJson(w http.ResponseWriter, httpCode int, payload string) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Gofiliate")
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.WriteHeader(httpCode)

	w.Write([]byte(payload))
	return
}
