package fiberfunc

import (
	"bre-api/handler"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func GetGeneric[T any](c *fiber.Ctx) error {
	var (
		id     int
		err    error
		result any
	)

	id, err = GetId(c)
	if err != nil {
		c.Status(ErrorConvertNumber.Startus)
		return c.SendString(ErrorConvertNumber.Err.Error())
	}

	result, err = handler.GetNoKey(new(T), id)
	if err != nil {
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}
	return c.JSON(result)
}

func GetGenericWithKey[T any](c *fiber.Ctx, joinTable []string) error {
	var (
		id     int
		err    error
		result any
	)

	id, err = GetId(c)
	if err != nil {
		c.Status(ErrorConvertNumber.Startus)
		return c.SendString(ErrorConvertNumber.Err.Error())
	}

	result, err = handler.GetAllWithKey(new(T), joinTable, id)
	if err != nil {
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}
	return OutJson(c, result, ErrorNil)
}

func GetAllGeneric[T any](c *fiber.Ctx) error {
	var (
		limit  int
		err    error
		result any
	)

	limit = Limit(c.Params("limit"))
	result, err = handler.GetNoKey(new([]T), 0, limit)
	if err != nil {
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}
	return OutJson(c, result, ErrorNil)
}
func GetAllGenericWithKey[T any](c *fiber.Ctx, joinTable []string) error {
	var (
		limit  int
		err    error
		result any
	)

	limit = Limit(c.Params("limit"))
	result, err = handler.GetAllWithKey(new(T), joinTable, 1, limit)
	if err != nil {
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}
	return OutJson(c, result, ErrorNil)
}

// func UpdateGeneric[T any](c *fiber.Ctx) error {
// 	id, err := GetId(c)
// 	if err != nil {
// 		return OutJson(c, nil, ErrorConvertNumber)
// 	}
// 	var data = new(T)
// 	if err := c.BodyParser(&data); err != nil {
// 		return OutJson(c, nil, ErrorConvertJson)
// 	}
// 	reflect.ValueOf(data).Elem().FieldByName("ID").SetUint(uint64(id))
//
//  result, err := handler.Updata(&data)
// 	if err != nil {
// 		return OutJson(c, result, ErrorToSql)
// 	}
// 	return OutJson(c, result, ErrorNil)
// }

func UpdateGeneric[T any](c *fiber.Ctx) error {
	var (
		id     int
		err    error
		result any
		data   T
	)

	id, err = GetId(c)
	if err != nil {
		c.Status(ErrorConvertNumber.Startus)
		return c.SendString(ErrorConvertNumber.Err.Error())
	}

	if err = c.BodyParser(&data); err != nil {
		c.Status(ErrorConvertJson.Startus)
		return c.SendString(ErrorConvertJson.Err.Error())
	}
	reflect.ValueOf(data).Elem().FieldByName("ID").SetUint(uint64(id))

	result, err = handler.Updata(&data)
	if err != nil {
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}
	return OutJson(c, result, ErrorNil)
}

func CreateGeneric[T any](c *fiber.Ctx) error {
	var (
		err    error
		result any
		data   T
	)

	err = c.BodyParser(&data)
	if err != nil {
		c.Status(ErrorConvertJson.Startus)
		return c.SendString(ErrorConvertJson.Err.Error())
	}

	result, err = handler.Create(&data)
	if err != nil {
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}
	return OutJson(c, result, ErrorNil)
}

func DeteleGeneric[T any](c *fiber.Ctx) error {
	var (
		id     int
		err    error
		result any
	)

	id, err = GetId(c)
	if err != nil {
		c.Status(ErrorConvertNumber.Startus)
		return c.SendString(ErrorConvertNumber.Err.Error())
	}

	result, err = handler.Delete(new(T), id)
	if err != nil {
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}
	return OutJson(c, result, ErrorNil)
}
