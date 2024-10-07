package service

import (
	"github.com/Fi44er/ton-backend/database"
	"github.com/Fi44er/ton-backend/dto"
	"github.com/Fi44er/ton-backend/model"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Create(ctx *fiber.Ctx, dto *dto.CreateTransaction) error {
	transaction := model.TransactionData{
		UserWalletAddress: dto.UserWalletAddress,
		DepositeDate:      dto.DepositeDate,
		ReceivingDate:     dto.ReceivingDate,
		Amount:            dto.Amount,
		Revards:           dto.Revards,
	}

	if err := database.DB.Db.Create(&transaction).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return ctx.Status(200).JSON(transaction)
}

func GetByWallet(ctx *fiber.Ctx, wallet string) error {
	transaction := new(model.TransactionData)

	if err := database.DB.Db.Where("user_wallet_address = ?", wallet).First(transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(fiber.Map{"error": "Транзакция не найдена"})
		}
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(200).JSON(transaction)
}

func GetAll(ctx *fiber.Ctx) error {
	var transactions []model.TransactionData

	if err := database.DB.Db.Find(&transactions).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(transactions) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"error": "Ни одной транзакции не найдено"})

	}
	return ctx.Status(200).JSON(transactions)
}

func Delete(ctx *fiber.Ctx, id int) error {
	transaction := new(model.TransactionData)
	database.DB.Db.Delete(transaction, id)
	return ctx.Status(200).JSON("OK")
}
