package lite

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS:  -lmatrix_creator_hal
#include "bus.h"
#include "sensors.h"
*/
import (
	"C"
)

import (
	"log"
)

// Init starts the MATRIXIOBus and exit the program if it can't.
// Every HAL function relies on the bus.
func Init() {
	if bool(!C.busInit()) {
		log.Fatal("matrixio_bus not initialized!\nIs MATRIX HAL installed?")
	}

	// Initialize Sensors
	C.uv_init()
	C.imu_init()
	C.humidity_init()
	C.pressure_init()

	// Initialize Everloop
	// Initialize GPIO
}