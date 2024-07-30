package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func StandandResponsePayload(w http.ResponseWriter, httpCode int, payload map[string]interface{}) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Gofiliate")
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.WriteHeader(httpCode)

	response := new(PayloadResponse)
	response.Response_code = httpCode
	response.Payload = payload

	js, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}
	w.Write(js)
	return

}
