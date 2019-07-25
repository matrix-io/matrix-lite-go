package matrix

/*
#cgo CXXFLAGS: -std=c++11
#cgo LDFLAGS:  -lmatrix_creator_hal
#include "bus.h"
#include "everloop.h"
#include "gpio.h"
#include "sensors.h"
*/
import (
	"C"
)

import (
	"log"
)

// Init starts the MATRIXIOBus or exits the program if it can't.
// Every HAL function relies on the bus.
func Init() {
	if bool(!C.busInit()) {
		log.Fatal("matrixio_bus not initialized!\nIs MATRIX HAL installed?")
	}

	C.everloop_init()
	C.gpio_init()
	C.uv_init()
	C.imu_init()
	C.humidity_init()
	C.pressure_init()
}
