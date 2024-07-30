package orm

import (
	"database/sql" //go inbuilt SQL package
)

/*
Delete generic function to perform a database UPDATE based on supplied table and data structure

[Accepts]

	conn *sql.DB mysql database connection
	table string which table to target
	record T generics reference to the data structure thats going to be updated

[Returns]

	record T amended record after the create has happened. Contains structure specific id (so this function is free from differing structure id names)
	err error any generated errors
*/
func DeleteWhere[T any](conn *sql.DB, table string, example T, conditions string, values ...interface{}) (err error) {
	// Get database connection
	db, err := GetConnection(conn)
	if err != nil {
		return err
	}

	// Execute delete query with the specified condition
	results := db.Table(table).Where(conditions, values...).Delete(&example)

	if results.Error != nil {
		return results.Error
	}

	return
}
