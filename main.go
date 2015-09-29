package main

import (
	"io/ioutil"

	"github.com/rameshvarun/ups-tools/reader"
)

func main() {
	data, err := ioutil.ReadFile("test.ups")
	if err != nil {
		panic(err)
	}

	_, err = reader.ReadUPS(data)
	if err != nil {
		panic(err)
	}
}
