package fiberfuncFile

import (
	"bre-api/config"
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/handler"
	"os"
	"reflect"
	"strings"

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
			return c.JSON(fiber.Map{
				"Error": err,
			})
		}

		fiberfuncConfig.GetFilePath(dataValue, StringName)
		
		fieldName := strings.Replace(StringName, "Upload", "", 1)
		if field := dataValue.FieldByName(fieldName + "Name"); field.IsValid() {
			field.SetString(file.Filename)
		}
		if field := dataValue.FieldByName(fieldName + "Path"); field.IsValid() {
			field.SetString(filePath)
		}
		Process = append(Process, StringName+" : Upload successful")
	}

	_, err := handler.Create(data)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}

	return c.JSON(Process)
}

func UpdateFile[T any](c *fiber.Ctx, modelsUploadfileString []string) error {
	id, err := fiberfuncConfig.GetId(c)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorConvertNumber.Startus)
		return c.JSON(fiberfuncConfig.ErrorConvertNumber.Err)
	}

	data := new(T)
	val := reflect.ValueOf(data).Elem()
	fieldId := val.FieldByName("ID")

	if fieldId.IsValid() && fieldId.CanSet() && fieldId.Kind() == reflect.Uint64 {
		fieldId.SetUint(uint64(id))
	} else {
		c.Status(fiberfuncConfig.ErrorFieldValid.Startus)
		return c.JSON(fiberfuncConfig.ErrorFieldValid.Err)
	}

	for _, StringData := range modelsUploadfileString {

		fieldName := strings.Join(strings.Split(StringData, "Upload"), "")
		oldFilePath := val.FieldByName(fieldName).String()
		if oldFilePath != "" {
			os.Remove(oldFilePath)
		}

		fileNew, err := c.FormFile(StringData)
		if err != nil {
			c.Status(fiberfuncConfig.ErrorConvertJson.Startus)
			return c.JSON(fiberfuncConfig.ErrorConvertJson.Err)
		}

		if fileNew.Size > config.MaxSizeFile {
			c.Status(fiberfuncConfig.ErrorFileSize.Startus)
			return c.JSON(fiberfuncConfig.ErrorFileSize.Err)
		}

		filePathNew := "./Uploadfile/" + fileNew.Filename
		err = c.SaveFile(fileNew, filePathNew)
		if err != nil {
			c.Status(fiberfuncConfig.ErrorToUploadFile.Startus)
			return c.JSON(fiberfuncConfig.ErrorToUploadFile.Err)
		}

		if field := val.FieldByName(fieldName + "Name"); field.IsValid() {
			field.SetString(fileNew.Filename)
		}
		if field := val.FieldByName(fieldName + "Path"); field.IsValid() {
			field.SetString(filePathNew)
		}

	}

	result, err := handler.Updata(&data)
	if err != nil {
		c.Status(fiberfuncConfig.ErrorToSql.Startus)
		return c.JSON(fiberfuncConfig.ErrorToSql.Err)
	}

	return c.JSON(result)
}

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
