package config

import (
	"errors"
	"time"
)

const (
	// config Sql not sqlite
	Host     = "Ip"
	Port     = 0
	User     = "Name"
	Password = "Password"
	Dbname   = "Name DataBase"
	// config Open Sql
	OpenServerDataBase    = true
	OpSql                 = false
	CrTable               = false
	CrUser                = false
	TLogH                 = false
	EndProcessOpen        = "NotProcessOpen"
	EndProcessCreateTable = "NotProcessOpen"
	EndProcessCreateData  = "NotProcessOpen"
	Limit                 = 10
	LoaclSqlite           = "grom/databases/testdb/dbreq.db"

	Openflber      = true
	MaxSizeFile    = 5 * 1024 * 1024
	CookieLifeTime = time.Hour * 72
)

var ErrorSqlLite = errors.New("Error Sql lite")
