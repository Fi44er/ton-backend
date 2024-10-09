package model

type Header struct {
	Id                int    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserWalletAddress string `json:"user_wallet_address" gorm:"type:string;not null"`
	Hash              string `json:"hash" gorm:"type:string;not null"`
	BodyID            int    `json:"body_id" gorm:"type:int;not null"`
	Body              Body   `json:"body" gorm:"foreignKey:BodyID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
