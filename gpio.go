package matrix

//#include "gpio.h"
import (
	"C"
)

import (
	"strings"
)

/////////////////////////////////
// GPIO BUS INIT & PIN CODES  //
///////////////////////////////
var setting map[string]int

func gpioInit() Gpio {
	C.gpio_init()

	// GPIO Pin Settings
	setting := make(map[string]int)

	// Mode
	setting["input"] = 0
	setting["output"] = 1

	// Function
	setting["digital"] = 0
	setting["pwm"] = 1

	// Digital Setting 
	setting["off"] = 0
	setting["on"] = 1

	return Gpio {}
}

/////////////////////////////////
//   GPIO STRUCT & METHODS    //
///////////////////////////////

// Gpio is used as the entrance point to communicate with
// the MATRIX GPIO Pinout. 
type Gpio struct{}

// SetFunction sets the pin to do Digital or PWM
func (*Gpio) SetFunction(pin int, function string){
	C.setFunction(C.int(pin), C.int(setting[strings.ToLower(function)]))
}

// GetDigital Reads a Pin's ON(1) or OFF(0) state (digital)
func (*Gpio) GetDigital(pin int) int{
	return int(C.getValue(C.int(pin)))
}

// SetMode sets the pin to input or output signal (digital)
func (*Gpio) SetMode(pin int, mode string){
	C.setMode(C.int(pin), C.int(setting[strings.ToLower(mode)]))
}

// SetDigital configures a pin to use Digital and output an ON/OFF signal
func (*Gpio) SetDigital(pin int, value string){
	C.setDigital(C.int(pin), C.int(setting[strings.ToLower(value)]))
}

// SetPWM configures a pin to use Pulse Width Modulation
func (*Gpio) SetPWM(pin int, percentage float32, frequency float32){
	C.setPWM(C.int(pin), C.float(percentage), C.float(frequency))
}

// SetServoAngle offers better PWM controls to manipulate a servo.
// Note: minPulseMs is a minimum pulse width for a PWM wave
func (*Gpio) SetServoAngle(pin int, angle float32, minPulseMs float32){
	C.setServoAngle(C.int(pin), C.float(angle), C.float(minPulseMs))
}