package fiberfuncFile

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/grom/handler"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func UploadfileGeneric[T any](c *fiber.Ctx, modelsUploadfileString []string) error {
	var data T
	dataValue := reflect.ValueOf(data).Elem()
	var Process []string
	for _, StringName := range modelsUploadfileString {
		file, filePath, err := fiberfuncConfig.UploadFile(c, StringName)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{"Error": err})
		}

		fiberfuncConfig.GetFilePath(dataValue, StringName)
		fiberfuncConfig.CheckCreateFile(StringName, dataValue, *file, filePath)
		Process = append(Process, StringName+" : Upload successful")
	}

	_, err := handler.Create(data)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}
	return c.JSON(Process)
}
