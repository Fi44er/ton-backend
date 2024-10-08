package model

type Body struct {
	Id                int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserWalletAddress string `json:"user_wallet_address" gorm:"type:string;not null"`
	DepositeDate      string `json:"deposite_date" gorm:"type:string;not null"`
	ReceivingDate     string `json:"receiving_date" gorm:"type:string;not null"`
	Amount            int    `json:"amount" gorm:"type:int;not null"`
	Rewards           int    `json:"rewards" gorm:"type:int;not null"`
}

type Header struct {
	Id                int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserWalletAddress string `json:"user_wallet_address" gorm:"type:string;not null"`
	Hash              string `json:"hash" gorm:"type:string;not null"`
	BodyID            int    `json:"body_id" gorm:"type:int;not null"`
	Body              Body   `gorm:"foreignKey:BodyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
