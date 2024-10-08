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
	// Проверяем валидность транзакции по хешу
	err := utils.CheckValidTransaction(dto.Header.Hash)
	if err != nil {
		return ctx.Status(400).JSON(fiber.Map{"error": err})
	}

	// Создаем запись в таблице Body
	bodyTransaction := model.Body{
		UserWalletAddress: dto.Body.UserWalletAddress,
		DepositeDate:      dto.Body.DepositeDate,
		ReceivingDate:     dto.Body.ReceivingDate,
		Amount:            dto.Body.Amount,
		Revards:           dto.Body.Revards,
	}

	// Сохраняем запись в таблице Body и получаем её ID
	if err := database.DB.Db.Create(&bodyTransaction).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Создаем запись в таблице Header, используя ID из созданной записи Body
	headerTransaction := model.Header{
		UserWalletAddress: dto.Body.UserWalletAddress,
		Hash:              dto.Header.Hash,
		BodyID:            bodyTransaction.Id, // Устанавливаем связь с Body
	}

	// Сохраняем запись в таблице Header
	if err := database.DB.Db.Create(&headerTransaction).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Возвращаем созданные записи
	return ctx.Status(200).JSON(fiber.Map{"status": "OK"})
}

func GetByWallet(ctx *fiber.Ctx, wallet string) error {
	transaction := new(model.Header)

	if err := database.DB.Db.Preload("Body").Where("user_wallet_address = ?", wallet).First(transaction).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Status(404).JSON(fiber.Map{"error": "Транзакция не найдена"})
		}
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return ctx.Status(200).JSON(transaction)
}

func GetAll(ctx *fiber.Ctx) error {
	var transactions []model.Header

	if err := database.DB.Db.Preload("Body").Find(&transactions).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	if len(transactions) == 0 {
		return ctx.Status(404).JSON(fiber.Map{"error": "Ни одной транзакции не найдено"})

	}
	return ctx.Status(200).JSON(transactions)
}

func Delete(ctx *fiber.Ctx, id int) error {
	transaction := new(model.Header)
	database.DB.Db.Delete(transaction, id)
	return ctx.Status(200).JSON("OK")
}
