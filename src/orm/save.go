package orm

import (

	//go inbuilt SQL package
	"database/sql"

	//visual debugging
	_ "gorm.io/driver/mysql" //gorm mysql package
)

/*
Save generic function to perform either an insert or an update backed on the primary keys of the incoming struct

[Accepts]

	conn *sql.DB mysql database connection
	table string which table to target
	record T generics reference to the data structure thats going to be created

[Returns]

	record T amended record after the create has happened. Contains structure specific id (so this function is free from differing structure id names)
	err error any generated errors
*/
func Save[T any](conn *sql.DB, table string, record *T) error {

	db, err := GetConnection(conn) //get our connection based on our established database connection

	db.Table(table).Save(&record)

	return err
}
