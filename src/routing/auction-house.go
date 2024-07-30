package routing

import (
	"github.com/gorilla/mux"

	"factory51/tornapi/src/auctions"
)

func RouteAuctions(r *mux.Router) {

	r.HandleFunc("/auctions/save-listings", auctions.EndpointSaveAuction).Methods("POST")

	r.HandleFunc("/auctions/save-listings/", auctions.EndpointSaveAuction).Methods("POST")

	r.HandleFunc("/auctions/save-listings", auctions.EndpointSaveAuction).Methods("GET")

}
