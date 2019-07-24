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
var LedLength int = int(C.everloopLength)

// Pass led slice to C
func cLedSet(leds []C.led) {
	pointer := unsafe.Pointer(&leds[0])
	C.everloop_set((*C.led)(pointer))
}

// Led represents the RGBW value of an LED (0-255)
type Led struct {
	R, G, B, W uint8
}

// obtain an LED
type rgbw interface {
	value() Led
}

func (l Led) value() Led {
	// return C.led{
	// 	r: C.int(l.R),
	// 	g: C.int(l.G),
	// 	b: C.int(l.B),
	// 	w: C.int(l.W),
	// }
	return l
}

// LedSet individually sets each MATRIX LED based on array index
func LedSet(color rgbw) {
	everloop := make([]C.led, 0, LedLength)
	defer fmt.Println(everloop)
	// defer cLedSet(everloop)

	switch input := reflect.TypeOf(color); {
	// Set all LEDs to one color
	case input == reflect.TypeOf(Led{}):
		for i := 0; i < 20; i++ {
			fmt.Println("i: ", i)
			//everloop[i] = color.value()
		}

		fmt.Println("YOU GAVE AN LED", color.value())

	case input.Kind() == reflect.String:
		fmt.Println("YOU GAVE A String")

	// Set LEDs to multiple colors
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
func LedSetAll(led Led) {
	everloop := []C.led{}

	for i := 0; i < LedLength; i++ {
		everloop = append(everloop, C.led{
			r: C.int(led.R),
			g: C.int(led.G),
			b: C.int(led.B),
			w: C.int(led.W),
		})
	}

	cLedSet(everloop)
}
