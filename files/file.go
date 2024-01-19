package file

import (
	"log"
	"os"
)




func ReadFile(filename string)[]byte {

	fileByte, err := os.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}
	return fileByte
}


// This function is to read a file in chunks rather than all at once

func ReadFileChun(filename string)