package auctions

import (
	"factory51/tornapi/src/database"
	"factory51/tornapi/src/orm"
	"fmt"
)

func GetIncomingAuctions(auction_id int64) (auctions []IncomingAuction, err error) {

	auction_example := IncomingAuction{}

	auctions, err = orm.GetByWhere(database.ADPT, "auctions_incoming", auction_example, "auction_id = ?", fmt.Sprintf("%v", auction_id))

	if err != nil {
		return
	}

	return

}
