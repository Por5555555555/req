package handler

import (
	"bre-api/databases"
	"fmt"

	"gorm.io/gorm"
)

func GetNoKey(myStruct any, GetWich ...int) (any, error) {
	db := databases.GetConn()
	return getIdOrAll(myStruct, *db, GetWich...)
}

func GetAllWithKey(myStruct any, TableSerlect []string, GetWich ...int) (any, error) {
	db := databases.GetConn()
	for _, val := range TableSerlect {
		db = db.Preload(val)
	}
	return getIdOrAll(myStruct, *db, GetWich...)
}

func getIdOrAll(myStruct any, db gorm.DB, GetWich ...int) (any, error) {
	var err error
	switch len(GetWich) {
	case 2:
		err = db.Limit(GetWich[1]).Find(myStruct).Error
	default:
		err = db.Find(myStruct, GetWich[0]).Error
	}
	return myStruct, err
}

func Create(Data any) (any, error) {
	db := databases.GetConn()
	err := db.Create(Data).Error
	return "Ok To Create Bro", err
}

func Updata(myStruct interface{}) (any, error) {
	db := databases.GetConn()
	err := db.Save(myStruct).Error
	return "Ok To Update Bro", err
}

func Delete(myStruct interface{}, Id int) (any, error) {
	db := databases.GetConn()
	var ItStruct = myStruct
	fmt.Print(ItStruct)
	err := db.Unscoped().Delete(ItStruct, Id).Error
	return "Ok To Delete", err
}
