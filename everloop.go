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

/////////////////////////////////////
// TODO: Adjust colors in colornames
////////////////////////////////////

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

// toCLed converts Led to a C readable LED
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

func anythingToCLed(color interface{}) C.led {
	// Determine how to set LED colors
	switch input := reflect.TypeOf(color); {

	//* Use string for all
	case input.Kind() == reflect.String:
		return stringToCLed(color.(string))
		fmt.Println("you used string")

	//* Use Led{} for all
	case input == reflect.TypeOf(Led{}):
		return color.(Led).toCLed()
		fmt.Println("you used struct Led")

	default:
		fmt.Print("UNEXPECTED INPUT: ")
		fmt.Println(reflect.TypeOf(color).Kind())
	}

	return C.led{}
}

// LedSet individually sets each MATRIX LED based on array index
func LedSet(color interface{}) {
	fmt.Println(reflect.TypeOf(color).Kind())

	// Create LEDs to set
	everloop := make([]C.led, LedLength())
	defer cLedSet(everloop)
	defer fmt.Println("\n\n\n", everloop)

	// Determine how to set LED colors
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
	//* Use a slice/array of the previous inputs
	case input.Kind() == reflect.Array || input.Kind() == reflect.Slice:
		list := reflect.ValueOf(color)

		for i := 0; i < list.Len(); i++ {
			everloop[i] = anythingToCLed(list.Index(i).Interface())
			// fmt.Println(list.Index(i))
		}

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
