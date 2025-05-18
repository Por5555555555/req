package fiberfuncCOUD

import (
	"bre-api/addOn/colortext"
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/gorm/handler"
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func UpdateGeneric[T any](c *fiber.Ctx) error {

	id, err := fiberfuncConfig.GetId(c)
	if err != nil {
		colortext.IsErr(err.Error() + "1")
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertNumber.Err)
	}

	input := new(T)
	if err := c.BodyParser(&input); err != nil {
		colortext.IsErr(err)
		c.Status(fiberfuncConfig.ErrorConvertJson.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertJson.Err)
	}

	oldData, err := handler.GetNoKey(new(T), id)
	if err != nil {
		colortext.IsErr(err)
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}

	newData := new(T)
	oldValue := reflect.ValueOf(oldData).Elem()
	newValue := reflect.ValueOf(newData).Elem()

	skipFields := map[string]bool{
		"ID":        true,
		"CreatedAt": true,
		"UpdatedAt": true,
		"DeletedAt": true,
	}

	for i := 0; i < oldValue.NumField(); i++ {
		fieldName := oldValue.Type().Field(i).Name
		if skipFields[fieldName] {
			continue
		}

		oldVal := oldValue.Field(i).Interface()
		newVal := reflect.ValueOf(input).Elem().Field(i).Interface()

		if oldVal != newVal {
			newValue.Field(i).Set(reflect.ValueOf(newVal))
			fmt.Println("Updated Field:", fieldName, "->", newVal)
		}
	}
	reflect.ValueOf(newData).Elem().FieldByName("ID").SetUint(uint64(id))

	jsonData, _ := json.MarshalIndent(newData, "", "  ")
	fmt.Println(string(jsonData))

	result, err := handler.Updata(&newData)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err.Error())
	}
	return c.JSON(fiber.Map{"Message": result})
}
