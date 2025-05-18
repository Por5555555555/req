package fiberfuncCOUD

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func DeteleGeneric[T any](c *fiber.Ctx) error {

	id, err := fiberfuncConfig.GetId(c)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertNumber.Err)
	}

	result, err := handler.Delete(new(T), id)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}
	return c.JSON(result)
}
