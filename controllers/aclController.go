package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/zhaizhonghao/auth/database"
	"github.com/zhaizhonghao/auth/models"
)

func GetAllEntries(c *fiber.Ctx) error {
	entries := []models.Entry{}
	database.DB.Find(&entries)
	fmt.Println(entries)
	return c.JSON(entries)
}

func AddACLEntry(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	entry := models.Entry{
		Email: data["email"],
	}

	database.DB.Create(&entry)

	return c.JSON(entry)
}

func DeleteACLEntry(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)

	if err != nil {
		return err
	}
	entry := models.Entry{
		Email: data["email"],
	}

	database.DB.Where("email=?", entry.Email).Delete(&entry)

	return c.JSON(entry)
}
