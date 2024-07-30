/*
file: goapi/cmd/import/database/database.go
description: database.go is the handler for setting up and managing the database connection
Author: david@gofiliate.com
Modified by: david:gofiliate.com
Last Modified: 2022-01-11
© 2022 Gotech Solutions Malta Limited (C72157), Trading as Gofiliate
*/
package database

import (
	"database/sql" //sql lib
	"fmt"          //used for string manipulation here

	_ "github.com/go-sql-driver/mysql"
)

const connection_skeleton = "%s:%s@tcp(%s:%v)/%s"

var ADPT *sql.DB //our database connection

func GetConnectionString(user string, pass string, host string, port int, database string) (connectionString string) {

	connectionString = fmt.Sprintf(connection_skeleton, user, pass, host, port, database)
	return
}

func Connect(connection_string string) (err error) {

	ADPT, err = sql.Open("mysql", connection_string)

	if err != nil {
		fmt.Printf("Error attempting to establish database connection: %s\n", err.Error())
	}

	err = ADPT.Ping()
	if err != nil {
		fmt.Printf("Error pinging database.ADPT: %v\n", err.Error())
	}
	return
}

func GetConnection() *sql.DB {

	return ADPT
}
