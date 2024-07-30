package auctions

type IncomingAuction struct {
	Id                   int64  `json:"incoming_id" gorm:"primaryKey;column:incoming_id"`
	Created              string `json:"created" gorm:"column:created"`
	UpdaterApiKey        string `json:"updater_api_key" gorm:"column:updater_api_key"`
	AuctionId            int64  `json:"auction_id" gorm:"column:auction_id"`
	SellerId             int64  `json:"seller_id" gorm:"seller_id"`
	SellerName           string `json:"seller_name" gorm:"column:seller_name"`
	BuyerId              int64  `json:"buyer_id" gorm:"column:buyer_id"`
	BuyerName            string `json:"buyer_name" gorm:"column:buyer_name"`
	ItemId               int64  `json:"item_id" gorm:"column:item_id"`
	ItemName             string `json:"item_name" gorm:"column:item_name"`
	Rarity               string `json:"rarity" gorm:"column:rarity"`
	CurrentNumBids       int    `json:"current_num_bids" gorm:"column:current_num_bids"`
	CurrentAuctionAmount int64  `json:"current_auction_amount" gorm:"column:current_auction_amount"`
	AuctionEnds          string `json:"auction_ends" gorm:"auction_ends"`
}
