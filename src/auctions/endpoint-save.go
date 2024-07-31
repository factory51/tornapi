package auctions

import (
	"encoding/json"
	"factory51/tornapi/src/database"
	"factory51/tornapi/src/orm"
	"factory51/tornapi/src/responses"
	"fmt"
	"net/http"
	"time"
)

func EndpointSaveAuction(w http.ResponseWriter, r *http.Request) {

	incoming_auction := IncomingAuction{}
	err := json.NewDecoder(r.Body).Decode(&incoming_auction)
	incoming_auction.Created = time.Now().Format("2006-01-02 15:04:05")

	if err != nil {
		msg := fmt.Sprintf("Error: Unable to decode supplied JSON. Please amend your request. : %v", err.Error())
		fmt.Printf("ERROR: %v\n", err.Error())
		responses.StandandResponseMessage(w, http.StatusBadRequest, msg) //trigger a bad request
		return
	}

	err = orm.Create(database.ADPT, "auctions_incoming", &incoming_auction)

	if err != nil {
		fmt.Printf("%v\n", err)
	}

	err = ProcessAuction(incoming_auction)

	if err != nil {
		fmt.Printf("[ERROR] %v\n", err.Error())
	}

}
