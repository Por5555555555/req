package fiberfunc

import (
	"bre-api/colortext"
	"bre-api/config"
	"reflect"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetId(c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return 0, ErrorConvertNumber
	}
	return id, nil
}

func GetUid(c *fiber.Ctx) (uint, error) {
	id, err := GetId(c)
	if err != nil {
		return 0, nil
	}
	return uint(id), nil
}

func OutText(c *fiber.Ctx, data any, err error) error {
	if err != nil {
		return c.Status(fiber.ErrBadRequest.Code).SendString(err.Error())
	}
	return c.JSON(data)
}

func unwarp(err error, Old int, New int) int {
	if err != nil {
		colortext.IsWanging("Unwarp For Number Convert")
		return New
	}
	return Old
}

func Limit(lim string) int {
	limit, err := strconv.Atoi(lim)
	if err == nil {
		return limit
	}
	return config.Limit
}

func getstring(Item []any) []string {
	var NameTable []string
	for _, val := range Item {
		NameTable = append(NameTable, reflect.TypeOf(val).Name())
	}
	return NameTable
}
