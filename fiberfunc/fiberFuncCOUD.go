package fiberfunc

import (
	"bre-api/config"
	"bre-api/handler"
	"os"
	"reflect"
	"strings"

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

func UploadfileGeneric[T any](c *fiber.Ctx, modelsUploadfileString []string) error {

	var data T
	dataValue := reflect.ValueOf(data).Elem()
	var Process []string
	for _, StringName := range modelsUploadfileString {
		file, err := c.FormFile(StringName)
		if err != nil {
			c.Status(ErrorConvertJson.Startus)
			return c.SendString(ErrorConvertJson.Err.Error())
		}
		if file.Size > config.MaxSizeFile {
			os.Remove(file.Filename)
			c.Status(ErrorFileSize.Startus)
			return c.SendString(ErrorFileSize.Err.Error())
		}

		filePath := "./Uploadfile/" + file.Filename
		err = c.SaveFile(file, filePath)
		if err != nil {
			c.Status(ErrorToUploadFile.Startus)
			return c.SendString(ErrorToUploadFile.Err.Error())
		}

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
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}

	return OutJson(c, Process, ErrorNil)
}

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

func UpdateFile[T any](c *fiber.Ctx, modelsUploadfileString []string) error {
	id, err := GetId(c)
	if err != nil {
		c.Status(ErrorConvertNumber.Startus)
		return c.SendString(ErrorConvertNumber.Err.Error())
	}

	data := new(T)
	val := reflect.ValueOf(data).Elem()
	fieldId := val.FieldByName("ID")

	if fieldId.IsValid() && fieldId.CanSet() && fieldId.Kind() == reflect.Uint64 {
		fieldId.SetUint(uint64(id))
	} else {
		c.Status(ErrorFieldValid.Startus)
		return c.SendString(ErrorFieldValid.Err.Error())
	}

	for _, StringData := range modelsUploadfileString {

		fieldName := strings.Join(strings.Split(StringData, "Upload"), "")
		oldFilePath := val.FieldByName(fieldName).String()
		if oldFilePath != "" {
			os.Remove(oldFilePath)
		}

		fileNew, err := c.FormFile(StringData)
		if err != nil {
			c.Status(ErrorConvertJson.Startus)
			return c.SendString(ErrorConvertJson.Err.Error())
		}

		if fileNew.Size > config.MaxSizeFile {
			os.Remove(fileNew.Filename)
			c.Status(ErrorFileSize.Startus)
			return c.SendString(ErrorFileSize.Err.Error())
		}

		filePathNew := "./Uploadfile/" + fileNew.Filename
		err = c.SaveFile(fileNew, filePathNew)
		if err != nil {
			c.Status(ErrorToUploadFile.Startus)
			return c.SendString(ErrorToUploadFile.Err.Error())
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
		c.Status(ErrorToSql.Startus)
		return c.SendString(ErrorToSql.Err.Error())
	}

	return OutJson(c, result, ErrorNil)
}
