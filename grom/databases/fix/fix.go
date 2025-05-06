package fix

import (
	"bre-api/databases"
	"errors"
	"fmt"
	"reflect"
)

func Fix[T any](G T) error {
	db := databases.GetConn()
	if db == nil {
        return errors.New("database connection is nil")
  }
	db.Exec("SET @count = 0;")

	gString, err := getfix(G)
	if err != nil {
		return err
	}

	fix := db.Exec(gString)
	if fix.Error != nil {
		return fix.Error
	}

	return nil
}

func getfix[T any](GString T) (string, error) {
  
	fixmap := map[string]string{
		"User":     "UPDATE users SET id = @count:=@count+1;",
		"Category": "UPDATE categorys SET id = @count:=@count+1;",
		"Money":    "UPDATE moneys SET id = @count:=@count+1;",
		"Provinc":  "UPDATE provinc SET id = @count:=@count+1;",
		"Source":   "UPDATE sources SET id = @count:=@count+1;",
		"Agency":   "UPDATE agencys SET id = @count:=@count+1;",
		"Auditor":  "UPDATE auditors SET id = @count:=@count+1;",
	  //"Money":    "UPDATE Money SET id = @count:=@count+1;",	
	}
  
	typeName := reflect.TypeOf(GString).Name()
	fmt.Println(typeName)
	fixSql, exists := fixmap[typeName]	
	if !exists {
		return fixSql, errors.New("Error find index")
	}
	fmt.Println("Is ok bro")
	return fixSql, nil
}

//1 3
//1 2
