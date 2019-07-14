package main

import (
	"fmt"

	"github.com/Hermitter/matrix-lite-go/hal"
)

func main() {
	fmt.Println("hello")
	hal.WrapTest()
	hal.BusInit()
}
