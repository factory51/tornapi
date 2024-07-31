package auctions

import (
	"fmt"

	"factory51/tornapi/src/database"
	"factory51/tornapi/src/orm"
)

func ProcessAuction(auction IncomingAuction) (err error) {

	exists, err := AuctionExists(auction.AuctionId)

	if err != nil {
		return
	}

	fmt.Printf("Auction: %v:%v\n", auction.AuctionId, exists)

	if !exists { //new

		new_auction := Auction{}
		new_auction.AuctionId = auction.AuctionId
		new_auction.ItemId = auction.ItemId
		new_auction.ItemName = auction.ItemName
		new_auction.Rarity = auction.Rarity
		new_auction.Bonus = auction.Bonus
		new_auction.Created = auction.Created
		new_auction.Ends = auction.AuctionEnds
		new_auction.SellerId = auction.SellerId
		new_auction.Ended = 0

		err = orm.Create(database.ADPT, "auctions", &new_auction) //add the auction to auctions table as first time we've seen it

		if err != nil {
			return
		}

	}

	auctions, err2 := GetIncomingAuctions(auction.AuctionId)

	if err2 != nil {
		err = err2
		return
	}

	for _, listing := range auctions { //loop through each of the entries for this auction_id

		exists, err3 := AuctionBidExists(listing.AuctionId, listing.CurrentNumBids)

		if err3 != nil {
			err = err3
			return
		}

		if !exists {

			bid := AuctionBid{}

			bid.AuctionId = listing.AuctionId
			bid.ItemId = listing.ItemId
			bid.ItemName = listing.ItemName
			bid.Created = listing.Created
			bid.BuyerId = listing.BuyerId
			bid.BidNum = listing.CurrentNumBids
			bid.Amount = listing.CurrentAuctionAmount
			err = orm.Create(database.ADPT, "auction_bids", &bid) //add the auction to auctions table as first time we've seen it

			if err != nil {
				return
			}

		}

	}

	return

}
