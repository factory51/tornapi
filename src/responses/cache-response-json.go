package responses

import (
	"net/http"
)

func CacheResponseJson(w http.ResponseWriter, r *http.Request, code int, payload string) {

	//fmt.Printf("is_cached: %v val: %v\n", is_cached, key)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Gofiliate")
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.WriteHeader(code)

	w.Write([]byte(payload))

}
