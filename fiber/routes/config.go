package routes

import (
	"bre-api/fiber/fiberfunc/fiberfuncConfig"
	"bre-api/grom/models"
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
	ConsiderKey = fiberfuncConfig.GetString([]any{
		Reqlacement{},
		Equipment{},
	})
	EquipmentKey = fiberfuncConfig.GetString([]any{
		Category{},
		Unit{},
	})
	FileKey = fiberfuncConfig.GetString([]any{
		Reqlacement{},
	})
	LocationKey = fiberfuncConfig.GetString([]any{
		Agency{},
		Province{},
	})
	NecessaryKey = fiberfuncConfig.GetString([]any{
		Location{},
	})
	PriceKey = fiberfuncConfig.GetString([]any{
		Equipment{},
	})
	ReqlacementKey = fiberfuncConfig.GetString([]any{
		Equipment{},
		Source{},
	})
)

var FileModelsUpload = []string{
	"UploadQuotation",
	"UploadPhoto",
	"UploadRepair",
	"UploadDiagram",
	"UploadDetails",
}

var FileIdConnect = "ReqlacementID"
