package lite

import (
	"github.com/Hermitter/matrix-lite-go/hal"
)

func Init() {
	hal.BusInit()
	//TODO: Initialize LEDs, Sensors, GPIOs
}

func ReadUV() float32 {
	return hal.ReadUV()
}
