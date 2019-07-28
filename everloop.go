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

///////////////////////////
//  EVERLOOP FUNCTIONS  //
/////////////////////////

func everloopInit() Led {
	C.everloop_init()
	return Led{
		Length: int(C.everloopLength),
	}
}

// cLedSet calls HAL to render the MATRIX Everloop.
func cLedSet(leds []C.led) {
	pointer := unsafe.Pointer(&leds[0])
	C.everloop_set((*C.led)(pointer))
}

// toCLed converts RGBW to a C readable LED.
// This is normally passed into cLedSet as an array
func (l RGBW) toCLed() C.led {
	return C.led{
		r: C.int(l.R),
		g: C.int(l.G),
		b: C.int(l.B),
		w: C.int(l.W),
	}
}

// toCLed converts any string or RGBW{} into a C.led
func toCLed(color interface{}) C.led {
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

	//* RGBW{}
	case input == reflect.TypeOf(RGBW{}):
		return color.(RGBW).toCLed()

	//* Invalid input
	default:
		return C.led{}
	}
}

///////////////////////////
//  STRUCTS & METHODS   //
/////////////////////////

// Led is used as the entrance point to communicate with
// the MATRIX everloop (That big ring of LEDs).
type Led struct {
	Length int
}

// RGBW represents the colors of a MATRIX LED (0-255)
type RGBW struct {
	R, G, B, W uint8
}

// Set parses any string, RGBW{}, or list of the latter
// to render the MATRIX everloop
func (led *Led) Set(color interface{}) error {
	everloop := make([]C.led, led.Length)
	defer cLedSet(everloop)

	// Determine how to set LED colors
	switch input := reflect.TypeOf(color); {

	//* Manually set each LED
	case input.Kind() == reflect.Array || input.Kind() == reflect.Slice:
		list := reflect.ValueOf(color)

		for i := 0; i < list.Len(); i++ {
			everloop[i] = toCLed(list.Index(i).Interface())
		}

	//* One color for all LEDs
	case input == reflect.TypeOf(RGBW{}) || input.Kind() == reflect.String:
		for i := 0; i < led.Length; i++ {
			everloop[i] = toCLed(color)
		}

	//* Invalid input
	default:
		return errors.New("LED configuration must be a string, RGBW{}, array, or slice")
	}

	return nil
}
