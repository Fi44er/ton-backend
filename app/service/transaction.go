package service

import (
	"github.com/Fi44er/ton-backend/database"
	"github.com/Fi44er/ton-backend/dto"
	"github.com/Fi44er/ton-backend/model"
	"github.com/Fi44er/ton-backend/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Create(ctx *fiber.Ctx, dto *dto.Req) error {
	err := utils.CheckValidTransaction(dto.Header.Hash)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err})
	}

	walletModel := new(model.Body)
	result := database.DB.Db.Where("user_wallet_address = ?", dto.Body.UserWalletAddress).First(walletModel)

	if result.RowsAffected == 0 {
		walletModel.UserWalletAddress = dto.Body.UserWalletAddress
		walletModel.DepositeDate = dto.Body.DepositeDate
		walletModel.ReceivingDate = dto.Body.ReceivingDate
		walletModel.Amount = dto.Body.Amount
		walletModel.Rewards = dto.Body.Rewards

		if err := database.DB.Db.Create(&walletModel).Error; err != nil {
			return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	} else if result.Error != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": result.Error.Error()})
	}

	headerTransaction := model.Header{
		UserWalletAddress: dto.Body.UserWalletAddress,
		Hash:              dto.Header.Hash,
		BodyID:            walletModel.Id,
	}

	if err := database.DB.Db.Create(&headerTransaction).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "OK"})
}

func GetByWallet(ctx *fiber.Ctx, wallet string) error {
	walletModel := new(model.Body)

	if err := database.DB.Db.Where("user_wallet_address = ?", wallet).First(walletModel).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(fiber.Map{"error": "Транзакция не найдена"})
		}
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(200).JSON(walletModel)
}

func GetAll(ctx *fiber.Ctx) error {
	var walletModel []model.Body

	if err := database.DB.Db.Preload("Headers").Find(&walletModel).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(walletModel) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"error": "Ни одной транзакции не найдено"})

	}
	return ctx.Status(200).JSON(walletModel)
}

func Delete(ctx *fiber.Ctx, id int) error {
	walletModel := new(model.Body)
	database.DB.Db.Delete(walletModel, id)
	return ctx.Status(200).JSON("OK")
}
