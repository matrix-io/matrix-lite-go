package hal

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS:  -lmatrix_creator_hal
#include "matrix.h"
#include "sensors.h"
*/
import (
	"C"
)

import (
	"log"
)

// BusInit will start the MATRIXIOBus and exit the program if it can't.
// Every HAL function relies on the bus.
func BusInit() {
	if bool(!C.busInit()) {
		log.Fatal("matrixio_bus not initialized!\nIs MATRIX HAL installed?")
	}
}

// UvRead ...
// func UvRead() {
// 	var num C.int = C.uvRead()
// 	fmt.Println(int(num))
// }
