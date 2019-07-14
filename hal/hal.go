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

	// Initialize Sensors
	C.uv_init()
	// Initialize Everloop
	// Initialize GPIO
}

//////////////////
// SENSOR READS \\

// ReadUV retrive data from MATRIX Creator UV Sensor
func ReadUV() float32 {
	return float32(C.uv_read().uv)
}

// func Parse (str string) *Cons {
// 	retCons := TranslateCCons2GoCons(cons_ptr)
// 	return retCons
// }

// func TranslateCCons2GoCons (c *C.cons_t) *Cons {
//     if c == nil {
//         return nil
//     }
//     return &Cons{
//         type: int(c.type),
//         car: TranslateCCons2GoCons(c.car),
//         cdr: TranslateCCons2GoCons(c.cdr),
//     }
// }
