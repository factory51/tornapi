package orm

import (

	//go inbuilt SQL package
	"database/sql"

	//visual debugging
	_ "gorm.io/driver/mysql" //gorm mysql package
)

/*
GetAll generic function to perform a database SELECT * FROM ... based on the supplied table name and example structure.

[Accepts]

	conn *sql.DB mysql database connection
	table string which table to target
	example T generics reference to the data structure thats going to be created

[Returns]

	object []T slice of supplied example T
	err error any generated errors
*/
func GetAllOrdered[T any](conn *sql.DB, table string, field string, order string, example T) (objects []T, err error) {

	db, err := GetConnection(conn) //get our connection based on our established database connection

	if err != nil {
		return
	}

	// Construct the order string
	orderString := field
	if order != "" {
		orderString += " " + order
	}

	results := db.Table(table).Order(orderString).Find(&objects)

	if results.Error != nil {
		err = results.Error
	}
	return

}
