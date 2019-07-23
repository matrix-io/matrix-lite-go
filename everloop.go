package matrix

//#include "everloop.h"
import (
	"C"
)

import (
	"unsafe"
)

// Led represents the RGBW value of an LED (0-255)
type Led struct {
	R, G, B, W uint8
}

// Led config that's passed to C
type cLed struct {
	Led C.led
}

// LedCount returns the number of LEDs on your MATRIX device
func LedCount() int {
	return int(C.everloopLength)
}

// Pass led slice to C
func cLedSet(leds []cLed) {
	pointer := unsafe.Pointer(&leds[0])
	C.everloop_set((*C.led)(pointer))
}

// LedSet individually sets each MATRIX LED based on array index
func LedSet(led []Led) {
	//TODO
}

// LedSetAll sets all MATRIX LEDs to one value
func LedSetAll(led Led) {
	everloop := []cLed{}

	for i := 0; i < LedCount(); i++ {
		everloop = append(everloop, cLed{Led: C.led{
			r: C.int(led.R),
			g: C.int(led.G),
			b: C.int(led.B),
			w: C.int(led.W),
		}})
	}

	cLedSet(everloop)
}
