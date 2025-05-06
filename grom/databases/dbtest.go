package databases

import (
	"bre-api/addOn/colortext"
	"bre-api/grom/models"
	"fmt"
	"reflect"
	"time"
)

func TestFunc(T time.Time) (string, error) {
	endProcess, err := testAll(T)
	if err != nil {
		return "", err
	}
	return endProcess, nil
}
func createModel(data []any) error {
	db := GetConn()
	testModelsSlice := data
	for _, models := range testModelsSlice {
		if err := db.Create(models).Error; err != nil {
			return err
		}
	}
	return nil
}

func deteleModel(data []any) error {
	db := GetConn()
	testModelsSlice := data
	for _, models := range testModelsSlice {
		if err := db.Unscoped().Delete(models).Error; err != nil {
			return err
		}
	}
	return nil
}

func processEnd(T time.Time) string {
	return time.Since(T).String()
}

func testAll(T time.Time) (string, error) {
	db := GetConn()

	var endProcess string
	TsCat := &models.Category{Name: "TsModels"}
	TsUnt := &models.Unit{Name: "TsUnit"}
	TsPro := &models.Province{Name: "TsProvince"}
	TsAge := &models.Agency{Name: "TsAgency"}
	TsMon := &models.Money{Name: "TsMoney"}
	TsSou := &models.Source{Name: "TsSource"}
	TsAud := &models.Auditor{
		FirstName: "TsAud",
		LastName:  "TsLast",
	}

	testModelsSlice := []any{
		TsCat,
		TsUnt,
		TsPro,
		TsAge,
		TsMon,
		TsSou,
		TsAud,
	}
	if err := createModel(testModelsSlice); err != nil {
		return processEnd(T), err
	}

	//-------------------------------------------------------------------------//

	TsEqu := &models.Equipment{
		Name:       "TsEqu",
		CategoryID: TsCat.ID,
		UnitID:     TsUnt.ID,
	}
	if err := db.Create(TsEqu).Error; err != nil {
		return processEnd(T), err
	}
	testModelsSlice = append(testModelsSlice, TsEqu)

	//-------------------------------------------------------------------------//

	TsReq := &models.Reqlacement{
		EquipmentID: TsEqu.ID,
		SourceID:    TsSou.ID,
		MoneyID:     TsMon.ID,
		Day:         time.Now(),
	}
	if err := db.Create(TsReq).Error; err != nil {
		return processEnd(T), err
	}
	testModelsSlice = append(testModelsSlice, TsReq)

	//-------------------------------------------------------------------------//

	TsLoc := &models.Location{
		Name:       "TsLoc",
		AgencyID:   TsAge.ID,
		ProvinceID: TsPro.ID,
	}
	if err := db.Create(TsLoc).Error; err != nil {
		return processEnd(T), err
	}
	testModelsSlice = append(testModelsSlice, TsLoc)

	//-------------------------------------------------------------------------//

	TsNec := &models.Necessary{
		LocationID: TsLoc.ID,
		Frequency:  "",
		Province:   "",
	}
	if err := db.Create(TsNec).Error; err != nil {
		return processEnd(T), err
	}
	testModelsSlice = append(testModelsSlice, TsNec)

	//-------------------------------------------------------------------------//

	TsPri := &models.Price{
		EquipmentID: TsEqu.ID,
		InMarket:    500,
		OutMarket:   700,
	}
	if err := db.Create(TsPri).Error; err != nil {
		return processEnd(T), err
	}
	testModelsSlice = append(testModelsSlice, TsPri)

	//-------------------------------------------------------------------------//

	TsCon := &models.Consider{
		ReqlacementID: TsReq.ID,
		EquipmentID:   TsEqu.ID,
		CanUse:        5,
		Missing:       5,
		All:           10,
	}
	if err := db.Create(TsCon).Error; err != nil {
		return processEnd(T), err
	}
	testModelsSlice = append(testModelsSlice, TsCon)

	//-------------------------------------------------------------------------//

	// TsFile := &models.File{
	// 	ReqlacementID: TsReq.ID,
	// 	Quotation:     "/Tq",
	// 	Photo:         "/TPh",
	// 	Repair:        "/TRep",
	// 	Diagram:       "/TDia",
	// 	Details:       "/TDet",
	// }
	// if err := db.Create(TsFile).Error; err != nil {
	// 	return processEnd(T), err
	// }
	// testModelsSlice = append(testModelsSlice, TsFile)

	//-------------------------------------------------------------------------//

	for _, models := range testModelsSlice {
		dataType := reflect.TypeOf(models)
		colortext.IsOk(dataType)
	}

	endProcess = processEnd(T)
	fmt.Println("To Create Data")
	fmt.Println("Enter to detele : ")
	// fmt.Scanf("Stop")
	// deteleModel(testModelsSlice)

	return endProcess, nil
}

func checkTy[T any](data T) string {
	chData := reflect.TypeOf(data).Name()
	return chData
}
