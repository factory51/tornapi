package helpers

import (
	"fmt"
	"os"
)

func LoadFile(file_path string) (data []byte) {

	data, err := os.ReadFile(file_path)

	if err != nil {
		fmt.Printf("[FATAL] Cannot load file [%v]. Reason: [%v]\n", file_path, err.Error())
		os.Exit(1) //quit here to help keep main clean
	}

	return

}
