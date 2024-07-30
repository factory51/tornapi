package helpers

import (
	"fmt"
	"os"

	"factory51/tornapi/src/config"
)

func LoadConfig(data []byte) {

	_, err := config.HandleConfig([]byte(data))

	if err != nil {
		fmt.Printf("[FATAL] Cannot load the raw config data into structure. Reason [%v]", err.Error()) //inform
		os.Exit(1)                                                                                     //and quit
	}

}
