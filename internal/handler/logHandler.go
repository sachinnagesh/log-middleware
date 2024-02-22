package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/sachinnagesh/log-middleware/cache"
	"github.com/sachinnagesh/log-middleware/internal/model"
)

func StoreLog(c *fiber.Ctx) error {

	logPayload := model.LogPayload{}
	/*err := json.Unmarshal(c.Request().Body(), &logPayload)
	if err != nil {
		fmt.Println("error while parsing request body" + err.Error())
		c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{"errors": err.Error()})
		return fiber.ErrBadRequest
	}*/

	if err := c.BodyParser(&logPayload); err != nil {
		//err := json.Unmarshal([]byte(content), &res1)
		fmt.Println("error while parsing request body" + err.Error())
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "error parsing request body!!!", "data": err.Error()})
	}
	err := cache.GetLogCache().AddLog(logPayload)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Error caching log!!!", "data": err.Error()})

	}

	return c.JSON(fiber.Map{"status": "success", "message": "log cached successfully!!!", "data": logPayload})

}

func GetLogsCount(c *fiber.Ctx) error {
	//c.SendString(strconv.Itoa(cache.GetLogCache().GetLogsCount()))
	return c.JSON(fiber.Map{
		"status": "success",
		"data":   map[string]int{"count": cache.GetLogCache().GetLogsCount()},
	})

}
