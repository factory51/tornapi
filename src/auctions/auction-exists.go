package auctions

import (
	"factory51/tornapi/src/database"
	"fmt"
)

func AuctionExists(auction_id int64) (exists bool, err error) {

	sql := "SELECT count(1) from auctions where auction_id = ?"

	query, err := database.ADPT.Prepare(sql)

	if err != nil {
		fmt.Printf("AuctionExists: Cannot Prepare: %v\n", err.Error())
	}

	rows, err := query.Query(auction_id)

	if err != nil {
		fmt.Printf("AuctionExists: Cannot Query: %v\n", err.Error())
	}

	for rows.Next() {

		var valid int
		err = rows.Scan(&valid)

		if err != nil {
			fmt.Printf("AuctionExists: Cannot Scan: %v\n", err.Error())
		}

		if valid == 1 {
			exists = true
		}
	}

	return

}
