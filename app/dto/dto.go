package dto

type Req struct {
	Header Header
	Body   Body
}

type Body struct {
	UserWalletAddress string `json:"user_wallet_address"`
	DepositeDate      string `json:"deposite_date"`
	ReceivingDate     string `json:"receiving_date"`
	Amount            int    `json:"amount"`
	Rewards           int    `json:"rewards"`
}

type Header struct {
	Hash string `json:"hash"`
}
