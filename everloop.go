package lite

//#include "everloop.h"
import (
	"C"
)

import (
	"fmt"
	"unsafe"
)

// LedCount returns the number of LEDs on your MATRIX device
func LedCount() int {
	return int(C.everloopLength)
}

// type Led struct{
// 	r uint8
// 	g uint8
// 	b uint8
// 	w uint8
// }

// CLed represents an LED color value for C++
type CLed struct {
	Led C.Led
}

// Pass led slice pointer to HAL
func ledSet(leds []CLed) {
	pointer := unsafe.Pointer(&leds[0])
	C.everloop_set((*C.Led)(pointer))
	// C.everloop_set(leds)
}

// DELETE THIS LATER JUST FOR TESTING
func TestLed() {
	leds := []CLed{}
	for i := 0; i < LedCount(); i++ {
		leds = append(leds, CLed{Led: C.Led{
			r: 0,
			g: 0,
			b: 0,
			w: 0,
		}})
	}

	ledSet(leds)
	fmt.Println(leds)
}
