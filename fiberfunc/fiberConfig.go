package fiberfunc

import (
	"bre-api/models"
	"errors"
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
	ErrorConvertJson   = errors.New("Error Convert Json to string")
	ErrorConvertNumber = errors.New("Error Convet String to number")
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
