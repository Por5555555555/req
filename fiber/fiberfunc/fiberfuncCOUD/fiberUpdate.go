package fiberfuncCOUD

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/handler"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func UpdateGeneric[T any](c *fiber.Ctx) error {

	id, err := fiberfuncConfig.GetId(c)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertNumber.Err)
	}

	data := new(T)
	if err := c.BodyParser(&data); err != nil {
		c.Status(fiberfuncConfig.ErrorConvertJson.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertJson.Err)
	}
	reflect.ValueOf(data).Elem().FieldByName("ID").SetUint(uint64(id))

	result, err := handler.Updata(&data)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}
	return c.JSON(result)
}
