package matrix

//#include "everloop.h"
import (
	"C"
)

import (
	"fmt"
	"reflect"
	"unsafe"

	"golang.org/x/image/colornames"
)

// LedLength returns the number of LEDs on your MATRIX device
func LedLength() int {
	return int(C.everloopLength)
}

// cLedSet renders MATRIX everloop
func cLedSet(leds []C.led) {
	pointer := unsafe.Pointer(&leds[0])
	C.everloop_set((*C.led)(pointer))
}

// Led represents the RGBW value of a MATRIX LED (0-255)
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

// stringToCled contains the following colors
// https://godoc.org/golang.org/x/image/colornames#pkg-variables
func stringToCLed(color string) C.led {
	led := colornames.Map[color]

	return C.led{
		r: C.int(led.R),
		g: C.int(led.G),
		b: C.int(led.B),
		w: C.int(0),
	}
}

// LedSet individually sets each MATRIX LED based on array index
func LedSet(color interface{}) {
	// Create LEDs to set
	everloop := make([]C.led, LedLength())
	defer cLedSet(everloop)

	// Determine how to set LEDs
	switch input := reflect.TypeOf(color); {

	//* Use Led{} for all
	case input == reflect.TypeOf(Led{}):
		for i := 0; i < LedLength(); i++ {
			everloop[i] = color.(Led).toCLed()
		}
	//* Use string for all
	case input.Kind() == reflect.String:
		for i := 0; i < LedLength(); i++ {
			everloop[i] = stringToCLed(color.(string))
		}

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
