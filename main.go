package main

import (
	"io/ioutil"
)

func main() {
	ioutil.Discard.Write(nil)
}
