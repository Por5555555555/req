package databases

import (
	"bre-api/addOn/colortext"
	"bre-api/gorm/models"
	"time"
)

func TestFunc(T time.Time) (string, error) {
	endProcess, err := create(T)
	if err != nil {
		return "", err
	}
	return endProcess, nil
}

func endTask(T time.Time) string {
	return time.Since(T).String()
}

func create(T time.Time) (string, error) {
	var endProcess string
	db := GetConn()

	// Create Agency
	age := models.Agency{Name: "Test Agency"}
	if err := db.Create(&age).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Agency")

	// Create Auditor
	aud := models.Auditor{FirstName: "John", LastName: "Doe"}
	if err := db.Create(&aud).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Auditor")

	// Create Category
	cat := models.Category{Name: "Test Category"}
	if err := db.Create(&cat).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Category")

	// Create money
	mon := models.Money{Name: "Test Money"}
	if err := db.Create(&mon).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Money")

	// Create Province
	pro := models.Province{Name: "Test Province"}
	if err := db.Create(&pro).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Province")

	// Create Source
	sur := models.Source{Name: "Test Source"}
	if err := db.Create(&sur).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Source")

	// Create unit
	unt := models.Unit{Name: "Test Unit"}
	if err := db.Create(&unt).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Unit")

	// Create user
	user := models.User{
		FirstName: "Jane",
		LastName:  "Doe",
		LevelUser: "Admin",
		Email:     "Wow@com",
	}
	if err := db.Create(&user).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create User")

	// Create equipment
	Equ := models.Equipment{
		Name:       "Test Equipment",
		UnitID:     unt.ID,
		Category:   cat,
		CategoryID: cat.ID,
	}
	if err := db.Create(&Equ).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Equipment")

	// Create location
	loc := models.Location{
		Name:       "Test Location",
		AgencyID:   age.ID,
		ProvinceID: pro.ID,
		Agency:     age,
		Province:   pro,
	}
	if err := db.Create(&loc).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Location")

	// Create necessary
	nec := models.Necessary{
		LocationID: loc.ID,
		Location:   loc,
	}
	if err := db.Create(&nec).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Necessary")

	// Create Price
	pri := models.Price{
		EquipmentID: Equ.ID,
		Equipment:   Equ,
		InMarket:    3000,
		OutMarket:   4000,
	}
	if err := db.Create(&pri).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Price")

	// Create Requests
	req := models.Requests{
		AgencyID:    age.ID,
		Agency:      age,
		LocationID:  loc.ID,
		Location:    loc,
		EquipmentID: Equ.ID,
		Equipment:   Equ,
		CategoryID:  cat.ID,
		Category:    cat,
		Replace:     "Test Replace",
		Request:     "Test Request",
		PriceID:     pri.ID,
		Price:       pri,
		AuditorID:   aud.ID,
		Auditor:     aud,
		UserID:      user.ID,
		User:        user,
		Remark:      "Test Remark",
		Date:        T,
		Necessary:   nec,
	}
	if err := db.Create(&req).Error; err != nil {
		return endTask(T), err
	}
	colortext.IsOk("Create Requests")

	endProcess = T.String()
	return endProcess, nil
}
