package matrix

//#include "everloop.h"
import (
	"C"
)

import (
	"errors"
	"reflect"
	"unsafe"

	"golang.org/x/image/colornames"
)

////////////////////////////////////////
// TODO: Adjust colors in colornames //
//////////////////////////////////////

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

// ToCLed can convert string or Led{}
// may add more in the future
func ToCLed(color interface{}) C.led {
	// Determine how to set LED colors
	switch input := reflect.TypeOf(color); {

	//* string
	case input.Kind() == reflect.String:
		led := colornames.Map[color.(string)]

		return C.led{
			r: C.int(led.R),
			g: C.int(led.G),
			b: C.int(led.B),
			w: C.int(0),
		}

	//* Led{}
	case input == reflect.TypeOf(Led{}):
		return color.(Led).toCLed()

	//* Invalid input
	default:
		return C.led{}
	}
}

// LedSet parses any string, Led{}, or list of the latter
// to render the MATRIX everloop
func LedSet(color interface{}) error {
	everloop := make([]C.led, LedLength())
	defer cLedSet(everloop)

	// Determine how to set LED colors
	switch input := reflect.TypeOf(color); {

	//* Manually set each LED
	case input.Kind() == reflect.Array || input.Kind() == reflect.Slice:
		list := reflect.ValueOf(color)

		for i := 0; i < list.Len(); i++ {
			everloop[i] = ToCLed(list.Index(i).Interface())
		}

	//* One color for all LEDs
	case input == reflect.TypeOf(Led{}) || input.Kind() == reflect.String:
		for i := 0; i < LedLength(); i++ {
			everloop[i] = ToCLed(color)
		}

	//* Invalid input
	default:
		return errors.New("LED configuration must be a string, Led{}, array, or slice")
	}

	return nil
}
