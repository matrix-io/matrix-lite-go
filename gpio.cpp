#include "matrix_hal/gpio_control.h"
#include "bus.hpp"
extern "C"{
    #include "gpio.h"
}

matrix_hal::GPIOControl gpio_control;


void gpio_init(){
    gpio_control.Setup(&bus);
}

// - Set GPIO Pin state (digital)
// Paramters: int pin (0-15), int mode (0=INPUT or 1=OUTPUT)
bool setPinMode(int pin, int mode){
    return gpio_control.SetMode(pin, mode);
}

// - Set GPIO Pin function
// Parameters: int pin (0-15), int function (0=DIGITAL or 1=PWM)
bool setPinFunction(int pin, int function){
    return gpio_control.SetFunction(pin, function);
}

// - Read GPIO Pin state (digital only)
int getPinValue(int pin){
    return gpio_control.GetGPIOValue(pin);
}

// - Read all GPIO Pin states (digital only)
// returned as a 16bit integer 
int getPinValues(){
    return gpio_control.GetGPIOValues();
}

// - Set GPIO Pin Digital
bool setDigital(int pin, int value){
    return gpio_control.SetGPIOValue(pin, value);
}

// - Set GPIO Pin PWM
bool setPWM(int pin, float percentage, float frequency){
    return gpio_control.SetPWM(frequency, percentage, pin);
}

// - Set Servo angle through PWM
bool setServoAngle(int pin, float angle, float min_pulse_ms){
    return gpio_control.SetServoAngle(angle, min_pulse_ms, pin);
}