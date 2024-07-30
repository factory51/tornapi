package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StandandResponseMessage(w http.ResponseWriter, httpCode int, message string) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Gofiliate")
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.WriteHeader(httpCode)

	response := new(Response)
	response.Code = httpCode
	response.Message = message

	js, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}
	w.Write(js)
	return

}
