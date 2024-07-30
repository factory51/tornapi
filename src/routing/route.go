package routing

import (
	"fmt" //string formatting
	"net/http"

	// used for http control
	// gorilla handler functions to get around pesky CORS dumbfuckery
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux" //gorilla mux router for our routing.
)

func HandleRoutes(port int) {

	//our gorilla mux router
	router := mux.NewRouter()

	RouteAuctions(router)

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT"})

	//follwings is some useful output so you can see what url the api is being served on
	full_domain := fmt.Sprintf(":%v", port) //lock to a port not a url and let ngnix direct traffic

	//listen and serve biatches with our CORS exceptions added to the router
	err := http.ListenAndServe(full_domain, handlers.CORS(origins, headers, methods)(router))

	fmt.Printf("Routing Error: %v\n", err.Error())

}
