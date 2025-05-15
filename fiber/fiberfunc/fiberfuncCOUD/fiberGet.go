package fiberfuncCOUD

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func GetGeneric[T any](c *fiber.Ctx) error {

	id, err := fiberfuncConfig.GetId(c)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertJson.Err.Error())
	}

	result, err := handler.GetNoKey(new(T), id)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err.Error())
	}
	return c.JSON(result)
}

func GetAllGeneric[T any](c *fiber.Ctx) error {
	var (
		limit  int
		err    error
		result any
	)

	limit = fiberfuncConfig.Limit(c.Params("limit"))
	result, err = handler.GetNoKey(new([]T), 0, limit)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err.Error())
	}
	return c.JSON(result)
}

// For Key

func GetGenericWithKey[T any](c *fiber.Ctx, joinTable []string) error {
	var (
		id     int
		err    error
		result any
	)

	id, err = fiberfuncConfig.GetId(c)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertNumber.Err)
	}

	result, err = handler.GetAllWithKey(new(T), joinTable, id)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}
	return c.JSON(result)
}

func GetAllGenericWithKey[T any](c *fiber.Ctx, joinTable []string) error {
	var (
		limit  int
		err    error
		result any
	)

	limit = fiberfuncConfig.Limit(c.Params("limit"))
	result, err = handler.GetAllWithKey(new(T), joinTable, 1, limit)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err.Error())
	}
	return c.JSON(result)
}
