package main

import (
	"fmt"

	"github.com/matti/betterioutil"
)

func main() {
	contents, _ := betterioutil.ReadFileString("/etc/resolv.conf")
	fmt.Println(contents)
}
