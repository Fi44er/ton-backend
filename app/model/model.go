package model

import "time"

type TransactionData struct {
	Id                int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserWalletAddress string    `json:"user_wallet_address" gorm:"type:string;not null"`
	DepositeDate      time.Time `json:"deposite_date" gorm:"not null"`
	ReceivingDate     time.Time `json:"receiving_date" gorm:"not null"`
	Amount            int       `json:"amount" gorm:"type:int;not null"`
	Revards           int       `json:"revards" gorm:"type:int;not null"`
}
