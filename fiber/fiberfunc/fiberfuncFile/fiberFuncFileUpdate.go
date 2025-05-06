package fiberfuncFile

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/grom/handler"
	"os"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func UpdateFile[T any](c *fiber.Ctx, modelsUploadfileString []string) error {
	id, err := fiberfuncConfig.GetId(c)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertNumber.Err)
	}

	data := new(T)
	dataValue := reflect.ValueOf(data).Elem()
	fieldId := dataValue.FieldByName("ID")

	if fieldId.IsValid() && fieldId.CanSet() && fieldId.Kind() == reflect.Uint64 {
		fieldId.SetUint(uint64(id))
	} else {
		c.Status(fiberfuncConfig.ErrorFieldValid.Startus)
		return c.JSON(fiberfuncConfig.ErrorFieldValid.Err)
	}

	for _, StringData := range modelsUploadfileString {

		fieldName := strings.Join(strings.Split(StringData, "Upload"), "")
		oldFilePath := dataValue.FieldByName(fieldName).String()
		if oldFilePath != "" {
			os.Remove(oldFilePath)
		}

		fileNew, filePathNew, err := fiberfuncConfig.UploadFile(c, StringData)
		if err != nil {
			c.Status(fiber.StatusBadRequest)
			return c.JSON(fiber.Map{"Error": err})
		}

		fiberfuncConfig.CheckCreateFile(StringData, dataValue, *fileNew, filePathNew)
	}

	result, err := handler.Updata(&data)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}

	return c.JSON(result)
}
