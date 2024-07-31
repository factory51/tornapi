package auctions

import (
	"factory51/tornapi/src/database"
	"factory51/tornapi/src/orm"
	"fmt"
)

func GetAuctionBids(auction_id int64) (bids []AuctionBid, err error) {

	bid_example := AuctionBid{}

	bids, err = orm.GetByWhere(database.ADPT, "auction_bids", bid_example, "auction_id = ?", fmt.Sprintf("%v", auction_id))

	if err != nil {
		return
	}

	return

}
