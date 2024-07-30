package orm

import (

	//go inbuilt SQL package
	"database/sql"

	//visual debugging
	_ "gorm.io/driver/mysql" //gorm mysql package
)

/*
CountWhere generic function to perform a database COUNT * FROM WHERE ... with where and value being used to build the where clause.

[Accepts]

	conn *sql.DB mysql database connection
	table string which table to target
	example T generics reference to the data structure thats going to be created
	where string holds the basis of the where clause gorm will build up eg field_name = ?
	value string the value for gorm to replace into the ? holder in there where string

[Returns]

	object []T slice of supplied example T
	err error any generated errors
*/
func CountWhere[T any](conn *sql.DB, table string, example T, where string, value string) (count int64, err error) {

	db, err := GetConnection(conn) //get our connection based on our established database connection
	results := db.Table(table).Where(where, value).Count(&count)

	if results.Error != nil {
		err = results.Error
	}
	return

}

// easier to create a second to not have to play around with variadic params
func CountWhereDouble[T any](conn *sql.DB, table string, example T, where string, value1 string, value2 string) (count int64, err error) {

	db, err := GetConnection(conn) //get our connection based on our established database connection
	results := db.Table(table).Where(where, value1, value2).Count(&count)

	if results.Error != nil {
		err = results.Error
	}
	return

}

// still easier but now looking sloppy - if another one is needed just write the variadic func
func CountWhereTriplet[T any](conn *sql.DB, table string, example T, where string, value1 string, value2 string, value3 string) (count int64, err error) {

	db, err := GetConnection(conn) //get our connection based on our established database connection
	results := db.Table(table).Where(where, value1, value2, value3).Count(&count)

	if results.Error != nil {
		err = results.Error
	}
	return

}
