package dto

import (
	"time"

	"github.com/asaskevich/govalidator"
)

type CreateTransaction struct {
	UserWalletAddress string    `json:"user_wallet_address" valid:"required~Адрес кошелька обязателен к заполнению"`
	DepositeDate      time.Time `json:"deposite_date" valid:"required~Дата внесения депозита обязательна к заполнению"`
	ReceivingDate     time.Time `json:"receiving_date" valid:"required~Дата получения обязательна к заполнению"`
	Amount            int       `json:"amount" valid:"required~Сумма обязательна к заполнению"`
	Revards           int       `json:"revards" valid:"required~Награды обязательна к заполнению,int~Revards must be a number"`
}

func (t *CreateTransaction) Validate() error {
	_, err := govalidator.ValidateStruct(t)
	return err
}
