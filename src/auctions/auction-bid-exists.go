package auctions

import (
	"factory51/tornapi/src/database"
	"fmt"
)

func AuctionBidExists(auction_id int64, bid_num int) (exists bool, err error) {

	sql := "SELECT count(1) from auction_bids where auction_id = ? and bid_num = ?"

	query, err := database.ADPT.Prepare(sql)

	if err != nil {
		fmt.Printf("AuctionBidExists: Cannot Prepare: %v\n", err.Error())
	}

	rows, err := query.Query(auction_id, bid_num)

	if err != nil {
		fmt.Printf("AuctionBidExists: Cannot Query: %v\n", err.Error())
	}

	for rows.Next() {

		var valid int
		err = rows.Scan(&valid)

		if err != nil {
			fmt.Printf("AuctionBidExists: Cannot Scan: %v\n", err.Error())
		}

		if valid == 1 {
			exists = true
		}
	}

	return
}
