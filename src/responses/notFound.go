package responses

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NotFound(w http.ResponseWriter, route string) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Powered-By", "Gofiliate")
	w.Header().Set("X-Clacks-Overhead", "GNU Terry Pratchett")
	w.WriteHeader(http.StatusNotFound)

	response := new(ErrorResponse)
	response.Error_code = http.StatusNotFound
	response.Error_message = "The requested route " + route + " has not been found."

	js, err := json.Marshal(response)

	if err != nil {
		fmt.Println(err)
	}
	w.Write(js)
	return

}
