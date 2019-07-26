package matrix

//#include "gpio.h"
import (
	"C"
)

import (
	"errors"
	"strings"
)

// pinFunction interface configures a pin as PWM or Digital
// This also handles the ON/OFF mode of a pin
type pinFunction interface {
	pinFunction()
}

// Pwm configures a pin's PWM settings
type Pwm struct {
	Pin        int
	Percentage float32
	Frequency  float32
}

func (p Pwm) pinFunction() {
	C.setPinFunction(C.int(p.Pin), C.int(1)) // 1=PWM
	C.setPWM(C.int(p.Frequency), C.float(p.Percentage), C.float(p.Pin))
}

// Digital represents a pin's ON(1) & OFF(0) state
type Digital struct {
	Pin   int
	State string
}

func (d Digital) pinFunction() {
	var state int

	if strings.ToLower(d.State) == "ON" {
		state = 1
	} else if strings.ToLower(d.State) == "OFF" {
		state = 0
	}

	C.setPinFunction(C.int(d.Pin), C.int(0)) // 0=DIGITAL
	C.setDigital(C.int(d.Pin), C.int(state))
}

// SetPinMode configures a pin to be INPUT/OUTPUT.
// int pin (0-15), int mode ("INPUT" or "OUTPUT")
func SetPinMode(pin int, mode string) error {
	var setting int

	if strings.ToLower(mode) == "input" {
		setting = 0
	} else if strings.ToLower(mode) == "output" {
		setting = 1
	} else {
		return errors.New("String must be == \"INPUT\" || \"OUTPUT\"")
	}

	C.setPinMode(C.int(pin), C.int(setting))
	return nil
}

// ReadPin returns a 0 or 1 value to represent the on/off state of a pin.
// int pin (0-15)
func ReadPin(pin int) int {
	return int(C.getPinValue(C.int(pin)))
}

// SetPinFunc configures a pin to be DIGITAL/PWM.
// int pin (0-15), string function ("DIGITAL" or "PWM")
func SetPinFunc(conf pinFunction) {
	conf.pinFunction()
}

////////////////////////////////
// TODO: Review functions below
///////////////////////////////

// func SetPinFunc(pin int, function string) error {
// 	var setting int

// 	if strings.ToLower(function) == "digital" {
// 		setting = 0
// 	} else if strings.ToLower(function) == "pwm" {
// 		setting = 1
// 	} else {
// 		return errors.New("String must be == \"DIGITAL\" || \"PWM\"")
// 	}

// 	C.setPinFunction(C.int(pin), C.int(setting))
// 	return nil
// }

// // SetPinDigital
// func SetPinDigital(pin int, value string) error {
// 	var setting int

// 	if strings.ToLower(value) == "ON" {
// 		setting = 0
// 	} else if strings.ToLower(value) == "OFF" {
// 		setting = 1
// 	} else {
// 		return errors.New("String must be == \"ON\" || \"OFF\"")
// 	}

// 	C.setDigital(C.int(pin), C.int(setting))
// 	return nil
// }

// // SetPinPWM
// func SetPinPWM(setting Pwm) {
// 	C.setPWM(C.int(setting.Frequency), C.float(setting.Percentage), C.float(setting.Pin))
// }
