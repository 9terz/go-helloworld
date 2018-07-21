package main

import (
	"fmt"

	"github.com/9terz/go-helloworld/gotools"
	"github.com/9terz/go-helloworld/gotools/stringhelper"
)

func init() {
	fmt.Println("kuy")
}
func main() {
	fmt.Println("Hello World")
	fmt.Println(gotools.Add(10, 2))
	fmt.Println(stringhelper.Upper("DOG"))
}
