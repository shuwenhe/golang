package main

import (
	"io/ioutil"
)

func main() {
	_, err := ioutil.ReadDir("/abc")
	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
