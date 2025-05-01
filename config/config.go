package config

import "errors"

const (
	// config Sql not sqlite
	Host     = 0
	Port     = 0
	User     = 0
	Password = 0
	Dbname   = 0
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

	Openflber = true
)

var ErrorSqlLite = errors.New("Error Sql lite")
