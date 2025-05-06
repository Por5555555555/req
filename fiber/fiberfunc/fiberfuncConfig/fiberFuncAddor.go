package fiberfuncConfig

import (
	"bre-api/addOn/colortext"
	"bre-api/config"
	"errors"
	"fmt"
	"mime/multipart"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetId(c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt("id")
	return id, err
}

func GetUid(c *fiber.Ctx) (uint, error) {
	id, err := GetId(c)
	if err != nil {
		return 0, nil
	}
	return uint(id), nil
}

func OutJson(c *fiber.Ctx, data any, status ErrorHanderfliber) error {
	if status.Err != nil {
		return c.Status(status.Startus).SendString(status.Err.Error())
	}
	return c.JSON(data)
}

func Unwarp(err error, Old int, New int) int {
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

func GetString(Item []any) []string {
	var NameTable []string
	for _, val := range Item {
		NameTable = append(NameTable, reflect.TypeOf(val).Name())
	}
	return NameTable
}

func GetFilePath(val reflect.Value, fieldName string) (string, error) {
	fieldPath := val.FieldByName(fieldName)
	if !fieldPath.IsValid() {
		return "", errors.New("Invalid field")
	}
	filePath := fieldPath.String()
	if filePath == "" {
		return "", fmt.Errorf("File path not found")
	}

	return filePath, nil
}

func DeleteFileAll(filePathAll []string) error {
	for _, filePath := range filePathAll {
		if _, err := os.Stat(filePath); err != nil {
			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("Error to Detele File : ", filePath)
			}
		}
	}
	return nil
}

func UploadFile(c *fiber.Ctx, StringName string) (*multipart.FileHeader, string, error) {
	file, err := c.FormFile(StringName)
	if err != nil {
		return nil, "", fmt.Errorf("Error to Form File")
	}
	if file.Size > config.MaxSizeFile {
		return nil, "", fmt.Errorf("Error to big file")
	}

	filePath := "./Uploadfile/" + file.Filename
	err = c.SaveFile(file, filePath)
	if err != nil {
		return nil, "", fmt.Errorf("Error to filePath")
	}

	return file, filePath, nil
}

func CheckCreateFile(StringName string, dataValue reflect.Value, file multipart.FileHeader, filePath string) {
	fieldName := strings.Replace(StringName, "Upload", "", 1)
	if field := dataValue.FieldByName(fieldName + "Name"); field.IsValid() {
		field.SetString(file.Filename)
	}
	if field := dataValue.FieldByName(fieldName + "Path"); field.IsValid() {
		field.SetString(filePath)
	}
}
