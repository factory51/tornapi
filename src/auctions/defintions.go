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
	Bonus                string `json:"bonus" gorm:"column:bonus"`
	CurrentNumBids       int    `json:"current_num_bids" gorm:"column:current_num_bids"`
	CurrentAuctionAmount int64  `json:"current_auction_amount" gorm:"column:current_auction_amount"`
	AuctionEnds          string `json:"auction_ends" gorm:"auction_ends"`
}

type Auction struct {
	AuctionId int64  `json:"auction_id" gorm:"column:auction_id"`
	ItemId    int64  `json:"item_id" gorm:"column:item_id"`
	ItemName  string `json:"item_name" gorm:"column:item_name"`
	Rarity    string `json:"rarity" gorm:"column:rarity"`
	Bonus     string `json:"bonus" gorm:"column:bonus"`
	Created   string `json:"created" gorm:"column:created"`
	Ends      string `json:"ends" gorm:"column:ends"`
	SellerId  int64  `json:"seller_id" gorm:"column:seller_id"`
	Ended     int    `json:"ended" gorm:"column:ended"`
}

type AuctionBid struct {
	AuctionId int64  `json:"auction_id" gorm:"column:auction_id"`
	ItemId    int64  `json:"item_id" gorm:"column:item_id"`
	ItemName  string `json:"item_name" gorm:"column:item_name"`
	Created   string `json:"created" gorm:"column:created"`
	BuyerId   int64  `json:"buyer_id" gorm:"column:buyer_id"`
	BidNum    int    `json:"bid_num" gorm:"column:bid_num"`
	Amount    int64  `json:"amount" gorm:"column:amount"`
}
