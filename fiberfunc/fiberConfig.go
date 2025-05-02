package fiberfunc

import (
	"bre-api/models"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type (
	Agency      models.Agency
	Auditor     models.Auditor
	Category    models.Category
	Consider    models.Consider
	Equipment   models.Equipment
	File        models.File
	Location    models.Location
	Money       models.Money
	Necessary   models.Necessary
	Price       models.Price
	Province    models.Province
	Reqlacement models.Reqlacement
	Source      models.Source
	Unit        models.Unit
	User        models.User

	ErrorHanderfliber struct {
		Err     error
		Startus int
	}
)

var (
	ConsiderKey = getstring([]any{
		Reqlacement{},
		Equipment{},
	})
	EquipmentKey = getstring([]any{
		Category{},
		Unit{},
	})
	FileKey = getstring([]any{
		Reqlacement{},
	})
	LocationKey = getstring([]any{
		Agency{},
		Province{},
	})
	NecessaryKey = getstring([]any{
		Location{},
	})
	PriceKey = getstring([]any{
		Equipment{},
	})
	ReqlacementKey = getstring([]any{
		Equipment{},
		Source{},
	})
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
		Err:     errors.New("Error To Sqk"),
		Startus: fiber.StatusBadRequest,
	}
	ErrorNil = ErrorHanderfliber{
		Err: nil,
	}
)

var FileModelsUpload = []string{
	"UploadQuotation",
	"UploadPhoto",
	"UploadRepair",
	"UploadDiagram",
	"UploadDetails",
}

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
