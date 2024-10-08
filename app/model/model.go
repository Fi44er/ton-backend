package model

import "time"

type Header struct {
	Id                int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserWalletAddress string    `json:"user_wallet_address" gorm:"type:string;not null"`
	DepositeDate      time.Time `json:"deposite_date" gorm:"not null"`
	ReceivingDate     time.Time `json:"receiving_date" gorm:"not null"`
	Amount            int       `json:"amount" gorm:"type:int;not null"`
	Revards           int       `json:"revards" gorm:"type:int;not null"`
}

type Body struct {
	Id                int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserWalletAddress string `json:"user_wallet_address" gorm:"type:string;not null"`
	Boc               string `json:"boc" gorm:"type:string;not null"`
	Marker            string `json:"marker" gorm:"type:string;not null"`
}
