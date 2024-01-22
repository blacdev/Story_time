package file

import (
	"os"
)




func ReadFile(filename string)([]byte, error ){

	fileByte, err := os.ReadFile(filename)

	return fileByte, err
}


// This function is to read a file in chunks rather than all at once

func ReadFileChun(filename string){
	// read file in chunks
}