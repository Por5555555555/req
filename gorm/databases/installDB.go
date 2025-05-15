package databases

import (
	"bre-api/gorm/models"
)

func CreateTabel() error {
	myStruct := []interface{}{

		&models.Equipment{},
		&models.Category{},
		&models.Unit{},
		&models.Price{},

		&models.Location{},
		&models.Province{},
		&models.Agency{},
		&models.Auditor{},

		&models.Reqlacement{},
		&models.Money{},
		&models.Source{},
		&models.Consider{},
		&models.Necessary{},
		&models.File{},
		&models.Requests{},
	}

	err := DB.AutoMigrate(myStruct...)
	if err != nil {
		return err
	}
	return nil
}
