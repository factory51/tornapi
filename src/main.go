package main

import (
	"factory51/tornapi/src/helpers"
	"factory51/tornapi/src/routing"
	"flag"
	"fmt"
)

var config_file string //console flag from the path to the config file to load to bootstrap the application
var port int           //Console flag for overriding listening port
var mode string        //Console flag to set how app is running. production|development (productions only writes to a file development to console as well)

func init() { // handle flags and connections to data sources

	flag.StringVar(&config_file, "config", "./conf/app.conf.json", "-config /path/to/config.file") //config cli param
	flag.IntVar(&port, "port", 8081, "-port {OVERRIDE_PORT}")                                      //port cli param
	flag.StringVar(&mode, "mode", "production", "-mode {production|development}")                  //mode
	flag.Parse()

	config_raw_json := helpers.LoadFile(config_file)

	helpers.LoadConfig(config_raw_json)

	helpers.ConnectToDatabase()

}

func main() {
	fmt.Printf("Running TORN API\n")

	routing.HandleRoutes(port) //load the routes and server
}
