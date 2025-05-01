package fiberfunc

import (
	"bre-api/handler"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func GetGeneric[T any](c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
		return OutText(c, nil, err)
	}

	data, err := handler.GetNoKey(new(T), id)
	return OutText(c, data, err)
}

func GetAllGeneric[T any](c *fiber.Ctx) error {
	limit := Limit(c.Params("limit"))
	data, err := handler.GetNoKey(new([]T), 0, limit)
	return OutText(c, data, err)
}

func UpdateGeneric[T any](c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
		return OutText(c, nil, err)
	}

	var data = new(T)
	if err := c.BodyParser(&data); err != nil {
		return OutText(c, nil, ErrorConvertJson)
	}
	reflect.ValueOf(data).Elem().FieldByName("ID").SetUint(uint64(id))
	updateData, err := handler.Updata(&data)
	return OutText(c, updateData, err)
}

func CreateGeneric[T any](c *fiber.Ctx) error {
	var data = new(T)
	if err := c.BodyParser(&data); err != nil {
		return OutText(c, nil, ErrorConvertJson)
	}

	process, err := handler.Create(&data)
	return OutText(c, process, err)
}

func DeteleGeneric[T any](c *fiber.Ctx) error {
	id, err := GetId(c)
	if err != nil {
		return OutText(c, nil, err)
	}

	process, err := handler.Delete(new(T), id)
	return OutText(c, process, err)
}

func GetGenericWithKey[T any](c *fiber.Ctx, TableConnect []string) error {
	id, err := GetId(c)
	if err != nil {
		return OutText(c, nil, ErrorConvertNumber)
	}
	fmt.Println(TableConnect)
	data, err := handler.GetAllWithKey(new(T), TableConnect, id)
	return OutText(c, data, err)
}

func GetAllGenericWithKey[T any](c *fiber.Ctx, TableConnect []string) error {
	limit := Limit(c.Params("limit"))
	data, err := handler.GetAllWithKey(new(T), TableConnect, 1, limit)
	return OutText(c, data, err)
}
