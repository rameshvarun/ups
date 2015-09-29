package main

import (
	"os"

	"github.com/rameshvarun/ups-tools/reader"
)

func main() {
	file, err := os.Open("test.ups")
	if err != nil {
		panic(err)
	}

	_, err = reader.ReadUPS(file)
	if err != nil {
		panic(err)
	}
}
