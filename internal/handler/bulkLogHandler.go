package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/sachinnagesh/log-middleware/internal/model"
)

func BulkLogProcessor(c *fiber.Ctx) error {
	log.Info("Received request for bulk logs post")

	logs := []model.LogPayload{}

	if err := c.BodyParser(&logs); err != nil {
		//err := json.Unmarshal([]byte(content), &res1)
		fmt.Println("error while parsing request body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "error parsing request body!!!", "data": err.Error()})
	}

	return c.JSON(fiber.Map{"status": "success"})

}
