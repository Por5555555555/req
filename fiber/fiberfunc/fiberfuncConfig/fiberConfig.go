package fiberfuncConfig

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type (
	ErrorHanderfliber struct {
		Err     error
		Startus int
	}
)

var (
	ErrorConvertJson = ErrorHanderfliber{
		Err:     errors.New("Error Convert Json to string"),
		Startus: fiber.StatusBadRequest,
	}
	ErrorConvertNumber = ErrorHanderfliber{
		Err:     errors.New("Error Convet String to number"),
		Startus: fiber.StatusBadRequest,
	}
	ErrorToUploadFile = ErrorHanderfliber{
		Err:     errors.New("Error To Upload file"),
		Startus: fiber.StatusBadRequest,
	}
	ErrorFileSize = ErrorHanderfliber{
		Err:     errors.New("Error To file size to big "),
		Startus: fiber.StatusRequestEntityTooLarge,
	}
	ErrorDataIsNil = ErrorHanderfliber{
		Err:     errors.New("Error in Sql Index Is nil"),
		Startus: fiber.StatusBadRequest,
	}
	ErrorFieldValid = ErrorHanderfliber{
		Err:     errors.New("Error Invalid field type or field does not exist"),
		Startus: fiber.StatusBadRequest,
	}
	ErrorToSql = ErrorHanderfliber{
		Err:     errors.New("Error To Sql"),
		Startus: fiber.StatusBadRequest,
	}
	ErrorToDeteleFile = ErrorHanderfliber{
		Err:     errors.New("Error To Detele file"),
		Startus: fiber.ErrBadRequest.Code,
	}
	ErrorToHashPassword = ErrorHanderfliber{
		Err:     errors.New("Sus"),
		Startus: fiber.StatusBadRequest,
	}
	ErrorToEnv = ErrorHanderfliber{
		Err:     errors.New("Error To Bad Config"),
		Startus: fiber.ErrBadGateway.Code,
	}
	ErrorNil = ErrorHanderfliber{
		Err: nil,
	}
)

// ---------------------------------------------------------------------------------------------------------------------//
// Models
// ---------------------------------------------------------------------------------------------------------------------//
// Create
// -------------------------------------------------------------------//
// Detele
// -------------------------------------------------------------------//
// GetAll and Get
// -------------------------------------------------------------------//
// Update
// -------------------------------------------------------------------//
