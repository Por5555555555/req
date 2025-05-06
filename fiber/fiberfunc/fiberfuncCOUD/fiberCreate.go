package fiberfuncCOUD

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/grom/handler"

	"github.com/gofiber/fiber/v2"
)

func CreateGeneric[T any](c *fiber.Ctx) error {

	data := new(T)
	err := c.BodyParser(&data)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertJson.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertJson.Err)
	}

	result, err := handler.Create(&data)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}
	return c.JSON(result)
}
