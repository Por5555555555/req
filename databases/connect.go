package databases

import (
	"bre-api/colortext"
	"bre-api/config"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

var (
	user                  = config.User
	password              = config.Password
	host                  = config.Host
	port                  = config.Port
	dbname                = config.Dbname
	endProcessOpen        = config.EndProcessOpen
	endProcessCreateTable = config.EndProcessCreateTable
	endProcessCreateData  = config.EndProcessCreateData
	opSql                 = config.OpSql
	crTable               = config.CrTable
	crUser                = config.CrUser
	TLogH                 = config.TLogH
	ErrorSqlLite          = config.ErrorSqlLite
)

func openShortForm(opSql bool, TLogH bool) error {
	process_OpSql := time.Now()
	db, err := openSql(opSql, TLogH)
	DB = db
	endProcessOpen = time.Since(process_OpSql).String()
	return err
}

func CreateTable(crTable bool) error {
	process_crTab := time.Now()
	if !crTable {
		return nil
	}
	err := createTabel()
	endProcessCreateTable = time.Since(process_crTab).String()
	return err
}

func CreateData(crUser bool) error {
	process_TestU := time.Now()
	if !crUser {
		return nil
	}

	endProcess, err := testFunc(process_TestU)
	endProcessCreateData = endProcess
	return err
}

func ConnectDataBase() error {

	if err := openShortForm(opSql, TLogH); err != nil {
		return err
	}

	if err := CreateTable(crTable); err != nil {
		return err
	}

	if err := CreateData(crUser); err != nil {
		return err
	}

	fmt.Println("Database migration completed!")
	colortext.IsTime("Time OpSql : " + endProcessOpen)
	colortext.IsTime("Time CreateTable : " + endProcessCreateTable)
	colortext.IsTime("Time CreateUser : " + endProcessCreateData)
	return nil
}

func GetConn() *gorm.DB {
	if DB == nil {
		panic("U cannot connect database")
	}
	return DB
}

func openSql(T bool, TLogH bool) (*gorm.DB, error) {
	newLogger := getLog(TLogH)
	dns := "testdb/dbreq.db"
	if T {
		dns = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			user,
			password,
			host,
			port,
			dbname)

		db, err := gorm.Open(mysql.Open(dns), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
			Logger:                                   newLogger,
		})

		return db, err
	}

	db, err := gorm.Open(sqlite.Open(dns), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   newLogger,
	})
	if err != nil {
		err = ErrorSqlLite
	}

	return db, err
}

func getLog(TLogH bool) logger.Interface {
	var conColor = false
	var logLevel = logger.Error
	var SlowThreshold = time.Second

	if TLogH {
		conColor = true
		logLevel = logger.Info
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: SlowThreshold,
			LogLevel:      logLevel,
			Colorful:      conColor,
		},
	)
	return newLogger
}

// var DBMySQLCIMS1 *gorm.DB

// // ConnectMySQLSource connects to MySQL database
// func ConnectToMySQLCIMS1(ip, sid, user, pass string, port int) {
// 	// Connect to MySQL database
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, ip, port, sid)
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to MySQL database: %v", err)
// 	}
// 	fmt.Println("Connected to MySQL database")

// 	// Get a reference to the underlying *sql.DB
// 	sqlDB, err := db.DB()
// 	if err != nil {
// 		log.Fatal("failed to get sql.DB instance:", err)
// 	}
// 	sqlDB.SetMaxIdleConns(40)
// 	sqlDB.SetMaxOpenConns(300)
// 	sqlDB.SetConnMaxLifetime(time.Hour)

// 	DBMySQLCIMS1 = db
// }

// func GetMySQLCIMS1() *gorm.DB {
// 	return DBMySQLCIMS1
// }

// func CloseMySQLCIMS1() {
// 	_, err := DBMySQLCIMS1.DB()
// 	if err != nil {
// 		log.Fatal("failed to get sql.DB instance:", err)
// 	}
// }
