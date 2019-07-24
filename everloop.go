package matrix

//#include "everloop.h"
import (
	"C"
)

import (
	"fmt"
	"reflect"
	"unsafe"
)

// LedLength returns the number of LEDs on your MATRIX device
func LedLength() int {
	return int(C.everloopLength)
}

// Set everloop LEDs
func cLedSet(leds []C.led) {
	pointer := unsafe.Pointer(&leds[0])
	C.everloop_set((*C.led)(pointer))
}

// Led represents the RGBW value of an LED (0-255)
type Led struct {
	R, G, B, W uint8
}

func (l Led) toCLed() C.led {
	return C.led{
		r: C.int(l.R),
		g: C.int(l.G),
		b: C.int(l.B),
		w: C.int(l.W),
	}
}

// TODO implement
func stringToCLed(string) C.led {
	return C.led{}
}

// LedSet individually sets each MATRIX LED based on array index
func LedSet(color interface{}) {
	// Hold and then set LED values
	everloop := make([]C.led, LedLength())
	defer cLedSet(everloop)

	// Determine how to set LEDs
	switch input := reflect.TypeOf(color); {

	// Assign Led{} colors to each LED
	case input == reflect.TypeOf(Led{}):
		for i := 0; i < LedLength(); i++ {
			everloop[i] = color.(Led).toCLed()
		}

	case input.Kind() == reflect.String:
		fmt.Println("YOU GAVE A STRING")
	case input.Kind() == reflect.Array:
		fmt.Println("YOU GAVE AN ARRAY")
	case input.Kind() == reflect.Slice:
		fmt.Println("YOU GAVE A SLICE")

	default:
		fmt.Print("UNEXPECTED INPUT: ")
		fmt.Println(reflect.TypeOf(color).Kind())
	}
}

// LedSetAll sets all MATRIX LEDs to one value
// func LedSetAll(led Led) {
// 	everloop := []C.led{}

// 	for i := 0; i < LedLength(); i++ {
// 		everloop = append(everloop, C.led{
// 			r: C.int(led.R),
// 			g: C.int(led.G),
// 			b: C.int(led.B),
// 			w: C.int(led.W),
// 		})
// 	}

// 	cLedSet(everloop)
// }
