package fiberfuncFile

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/gorm/handler"
	"bre-api/gorm/models"
	"reflect"

	"github.com/gofiber/fiber/v2"
)

func UploadfileGeneric[T any](c *fiber.Ctx, modelsUploadfileString []string, connectTable string) error {
	connectId, err := fiberfuncConfig.GetId(c)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertNumber.Err)
	}

	data := new(T)
	dataValue := reflect.ValueOf(data).Elem()
	if connectTable != "" {
		dataValue.FieldByName(connectTable).SetInt(int64(connectId))
	}

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

	_, err = handler.Create(data)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}
	return c.JSON(Process)
}

func uploadFile(c *fiber.Ctx) error {

	FileCreate := new(models.File)
	err := c.BodyParser(FileCreate)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertJson.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertJson.Err)
	}

	return c.JSON("")
}
