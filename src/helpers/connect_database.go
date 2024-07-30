package helpers

import (
	"fmt"
	"os"

	"factory51/tornapi/src/config"

	"factory51/tornapi/src/database"
)

func ConnectToDatabase() {

	user := config.GetConfig().Database.Username //get our username
	pass := config.GetConfig().Database.Password //which password
	host := config.GetConfig().Database.Host     //which hostname
	port := config.GetConfig().Database.Port     //which port
	name := config.GetConfig().Database.Database //which  database

	conn := database.GetConnectionString(user, pass, host, port, name) //create our connection string

	err := database.Connect(conn) //connect to the database

	if err != nil {
		fmt.Printf("[FATAL] Cannot connect to the database on [%v]. Reason: [%v]\n", conn, err.Error()) //inform
		os.Exit(1)                                                                                      //and quit
	}

}
