package fiberfuncFile

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/gorm/handler"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func DeleteFile[T any](c *fiber.Ctx, modelsUploadfileString []string) error {
	id, err := fiberfuncConfig.GetId(c)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertNumber.Err)
	}

	data, err := handler.GetNoKey(new(T), id)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}

	val := reflect.ValueOf(data).Elem()
	var filePathAll []string
	var filePathAllStatus []fiber.Map

	for _, stringField := range modelsUploadfileString {
		fieldName := strings.Replace(stringField, "Upload", "", 1) + "Path"
		filePath, err := fiberfuncConfig.GetFilePath(val, fieldName)

		status := "Ready to delete"
		if err != nil {
			status = err.Error()
		} else {
			filePathAll = append(filePathAll, filePath)
		}

		filePathAllStatus = append(filePathAllStatus, fiber.Map{
			"field":  fieldName,
			"status": status,
		})
	}

	if err := fiberfuncConfig.DeleteFileAll(filePathAll); err != nil {
		c.Status(fiberfuncConfig.ErrorToDeteleFile.Startus)
		return c.JSON(fiberfuncConfig.ErrorToDeteleFile.Err)
	}

	_, err = handler.Delete(new(T), id)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToDeteleFile.Err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Files deleted successfully",
		"details": filePathAllStatus,
	})
}
