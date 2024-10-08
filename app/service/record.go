package service

import (
	"github.com/Fi44er/ton-backend/database"
	"github.com/Fi44er/ton-backend/dto"
	"github.com/Fi44er/ton-backend/model"
	"github.com/gofiber/fiber/v2"
)

func Update(ctx *fiber.Ctx, req *dto.Record) error {
	var records []model.Record

	// Получаем все записи (или можно реализовать фильтрацию по ID, если нужно)
	if err := database.DB.Db.Find(&records).Error; err != nil {
		return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	// Проходим по всем записям и обновляем только измененные поля
	for i := range records {
		updateData := req.Data[i]
		if err := database.DB.Db.Model(&records[i]).Updates(&updateData).Error; err != nil {
			return ctx.Status(500).JSON(fiber.Map{"error": err.Error()})
		}
	}

	return ctx.Status(200).JSON(fiber.Map{"status": "OK"})
}
