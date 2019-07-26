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

// Matrix is the main object that allows communication with MATRIX HAL
type Matrix struct {
	Led
	Imu
	Uv
	Pressure
	Humidity
}

// Init starts the MATRIXIOBus or exits the program if it can't.
// Every HAL function relies on the bus.
func Init() Matrix {
	if bool(!C.busInit()) {
		log.Fatal("matrixio_bus not initialized!\nIs MATRIX HAL installed?")
	}

	C.gpio_init()

	return Matrix{
		Led:      everloopInit(),
		Imu:      imuInit(),
		Uv:       uvInit(),
		Humidity: humidityInit(),
		Pressure: pressureInit(),
	}
}
