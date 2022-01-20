package main

import (
	"fmt"
	"os"
)

const (
	DIR = "/"
)

func main() {
	b := IsExist(DIR)
	fmt.Println("b = ", b)
}

func IsExist(path string)bool {
	FileInfo, err := os.Stat(path)
	fmt.Println("FileInfo = ", FileInfo)

	if err != nil {
		os.IsExist(err)
	}
	return true
}
