package main

import (
	"io/ioutil"

	"github.com/rameshvarun/ups/reader"
	"github.com/rameshvarun/ups/writer"
)

func main() {
	data, err := ioutil.ReadFile("mother3.ups")
	if err != nil {
		panic(err)
	}

	patch, err := reader.ReadUPS(data)
	if err != nil {
		panic(err)
	}

	newdata := writer.WriteUPS(patch)
	ioutil.WriteFile("mother3.out", newdata, 0644)
}
