package orm

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
GetConnection - uses an established *sql.DB connection to allow gORM to communicate with the database.

[Accepts]

	conn *sql.DB database connection

[Returns]

	db *gorm.DB connection to query database
	err error any generated errors
*/
func GetConnection(conn *sql.DB) (db *gorm.DB, err error) {

	mode := logger.Silent

	db, err = gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{
		Logger: logger.Default.LogMode(mode),
	})

	return

}
