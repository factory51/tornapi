package responses

import (
	"net/http"
	"encoding/json"
	"fmt"
)


func NotImplemented(w http.ResponseWriter, route string) {

	
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Gofiliate")
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.WriteHeader(http.StatusNotImplemented)

	response := new(ErrorResponse)
	response.Error_code = http.StatusNotImplemented
	response.Error_message = "The requested route " + route + " has not been implemeneted yet. Please stand by."

	js, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}
	w.Write(js)		
	return

}