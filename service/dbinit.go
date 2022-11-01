package service

import (
	"database/sql"

	"github.com/mateenbagheri/mysmall-wallet/configs"
)


const setLimits bool = true

var Mysql *sql.DB = configs.ConnectMySQL(setLimits)
