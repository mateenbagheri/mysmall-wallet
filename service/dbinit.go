package service

import (
	"database/sql"

	"github.com/mateenbagheri/mysmall-wallet/configs"
)

var Mysql *sql.DB = configs.ConnectMySQL()
